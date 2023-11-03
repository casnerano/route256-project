package internal

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"route256/cart/pkg/logger"
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
	cancelUnpaidOrder worker
}

type depRepository struct {
	order repository.Order
	stock repository.Stock
}

type depService struct {
	order *order.Order
	stock *stock.Stock
}

type application struct {
	config        *config.Config
	server        *server.Server
	transactor    repository.Transactor
	statusSender  order.StatusSender
	depRepository *depRepository
	depService    *depService
	depWorker     *depWorker
	logger        *zap.Logger
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

	if app.config.Database.DSN == "" {
		return nil, fmt.Errorf("database dsn is required")
	}

	var pool *pgxpool.Pool
	pool, err = pgxpool.New(context.TODO(), app.config.Database.DSN)
	if err != nil {
		return nil, err
	}

	sqlTransactor := sqlstore.NewTransactor(pool)
	app.transactor = sqlTransactor

	app.statusSender, err = order.NewKafkaStatusSender(
		app.config.Order.StatusSender.Brokers,
		app.config.Order.StatusSender.Topic,
	)
	if err != nil {
		return nil, err
	}

	app.depRepository = &depRepository{
		order: sqlstore.NewOrderRepository(sqlTransactor),
		stock: sqlstore.NewStockRepository(sqlTransactor),
	}

	app.depService = &depService{
		order: order.New(
			app.statusSender,
			app.transactor,
			app.depRepository.order,
			app.depRepository.stock,
		),
		stock: stock.New(app.depRepository.stock),
	}

	app.depWorker = &depWorker{
		cancelUnpaidOrder: order.NewCancelUnpaidWorker(
			app.depService.order,
			time.Duration(app.config.Order.CancelUnpaidTimeout)*time.Second,
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
	a.server, err = server.New(a.config.Server, a.depService.order, a.depService.stock)

	return err
}

func (a *application) RunGRPCServer() error {
	return a.server.RunGRPC()
}

func (a *application) RunHTTPServer() error {
	return a.server.RunHTTP()
}

func (a *application) Shutdown() error {
	if err := a.statusSender.Close(); err != nil {
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
