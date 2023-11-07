package internal

import (
	"context"
	"fmt"
	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"route256/cart/pkg/logger"
	"route256/cart/pkg/trace"
	"route256/loms/internal/config"
	"route256/loms/internal/repository"
	"route256/loms/internal/repository/sqlstore"
	"route256/loms/internal/server"
	"route256/loms/internal/service/order"
	"route256/loms/internal/service/stock"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type worker interface {
	Run(ctx context.Context) error
}

type depWorker struct {
	cancelUnpaidOrder       worker
	senderOrderStatusOutbox worker
}

type depRepository struct {
	order             repository.Order
	stock             repository.Stock
	orderStatusOutbox repository.OrderStatusOutbox
}

type depService struct {
	order        *order.Order
	stock        *stock.Stock
	statusSender order.StatusSender
}

type application struct {
	config        *config.Config
	server        *server.Server
	transactor    repository.Transactor
	depRepository *depRepository
	depService    *depService
	depWorker     *depWorker
	logger        *zap.Logger
	traceProvider *sdktrace.TracerProvider
}

func NewApp() (*application, error) {
	var err error
	app := application{}

	app.config, err = config.New()
	if err != nil {
		return nil, err
	}

	app.logger, err = logger.New("loms")
	if err != nil {
		return nil, err
	}

	app.traceProvider, err = trace.NewTraceProvider("loms")
	if err != nil {
		return nil, err
	}

	if app.config.Database.DSN == "" {
		return nil, fmt.Errorf("database dsn is required")
	}

	var pool *pgxpool.Pool
	pgxConfig, err := pgxpool.ParseConfig(app.config.Database.DSN)
	if err != nil {
		return nil, err
	}

	pgxConfig.ConnConfig.Tracer = otelpgx.NewTracer()
	pgxConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeDescribeExec

	pool, err = pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		return nil, err
	}

	sqlTransactor := sqlstore.NewTransactor(pool)
	app.transactor = sqlTransactor

	app.depRepository = &depRepository{
		order:             sqlstore.NewOrderRepository(sqlTransactor),
		stock:             sqlstore.NewStockRepository(sqlTransactor),
		orderStatusOutbox: sqlstore.NewOrderStatusOutboxRepository(sqlTransactor),
	}

	statusSender, err := order.NewKafkaStatusSender(
		app.config.Order.StatusSender.Brokers,
		app.config.Order.StatusSender.Topic,
	)
	if err != nil {
		return nil, err
	}

	app.depService = &depService{
		order: order.New(
			app.transactor,
			app.depRepository.order,
			app.depRepository.stock,
			app.depRepository.orderStatusOutbox,
		),
		stock:        stock.New(app.depRepository.stock),
		statusSender: statusSender,
	}

	app.depWorker = &depWorker{
		cancelUnpaidOrder: order.NewCancelUnpaidWorker(
			app.depService.order,
			time.Duration(app.config.Order.CancelUnpaidTimeout)*time.Second,
			app.logger,
		),
		senderOrderStatusOutbox: order.NewStatusOutbox(
			app.depRepository.orderStatusOutbox,
			statusSender,
			app.logger,
		),
	}

	err = app.init()
	if err != nil {
		return nil, fmt.Errorf("failed init server: %w", err)
	}

	return &app, nil
}

func (a *application) init() error {
	var err error
	a.server, err = server.New(a.config.Server, a.depService.order, a.depService.stock, a.logger)

	return err
}

func (a *application) RunGRPCServer() error {
	return a.server.RunGRPC()
}

func (a *application) RunHTTPServer() error {
	return a.server.RunHTTP()
}

func (a *application) Shutdown() error {
	if err := a.traceProvider.Shutdown(context.Background()); err != nil {
		return err
	}

	if err := a.depService.statusSender.Close(); err != nil {
		return err
	}

	if err := a.server.ShutdownHTTP(); err != nil {
		return err
	}

	return a.server.ShutdownGRPC()
}

func (a *application) RunCancelUnpaidOrderWorker(ctx context.Context) error {
	return a.depWorker.cancelUnpaidOrder.Run(ctx)
}

func (a *application) RunOutboxWorker(ctx context.Context) error {
	return a.depWorker.senderOrderStatusOutbox.Run(ctx)
}
