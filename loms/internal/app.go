package internal

import (
	"context"
	"fmt"
	"route256/loms/internal/config"
	"route256/loms/internal/repository"
	"route256/loms/internal/repository/memstore"
	"route256/loms/internal/server"
	orderHandler "route256/loms/internal/server/handler/order"
	stockHandler "route256/loms/internal/server/handler/stock"
	"route256/loms/internal/service/order"
	"route256/loms/internal/service/stock"
	pbOrder "route256/loms/pkg/proto/order/v1"
	pbStock "route256/loms/pkg/proto/stock/v1"
	"time"
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
	depRepository *depRepository
	depService    *depService
	depWorker     *depWorker
}

func NewApp() (*application, error) {
	var err error
	app := application{}

	app.config, err = config.New()
	if err != nil {
		return nil, err
	}

	app.depRepository = &depRepository{
		order: memstore.NewOrderRepository(),
		stock: memstore.NewStockRepository(),
	}

	app.depService = &depService{
		order: order.New(app.depRepository.order, app.depRepository.stock),
		stock: stock.New(app.depRepository.stock),
	}

	app.depWorker = &depWorker{
		cancelUnpaidOrder: order.NewCancelUnpaidWorker(
			app.depService.order,
			time.Duration(app.config.Order.CancelUnpaidTimeout)*time.Second,
		),
	}

	err = app.initServer()
	if err != nil {
		return nil, fmt.Errorf("failed init server: %w", err)
	}

	return &app, nil
}

func (a *application) initServer() error {
	var err error

	a.server, err = server.New(a.config.Server)
	if err != nil {
		return err
	}

	pbOrder.RegisterOrderServer(a.server.GRPC, orderHandler.NewHandler(a.depService.order))
	pbStock.RegisterStockServer(a.server.GRPC, stockHandler.NewHandler(a.depService.stock))

	return nil
}

func (a *application) RunServer() error {
	return a.server.Run()
}

func (a *application) ShutdownServer() error {
	return a.server.Shutdown()
}

func (a *application) RunCancelUnpaidOrderWorker(ctx context.Context) error {
	return a.depWorker.cancelUnpaidOrder.Run(ctx)
}
