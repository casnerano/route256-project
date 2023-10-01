package internal

import (
	"context"
	"fmt"
	"route256/loms/internal/config"
	"route256/loms/internal/repository"
	"route256/loms/internal/repository/memstore"
	"route256/loms/internal/server"
	"route256/loms/internal/service/order"
	"route256/loms/internal/service/stock"
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
	if err := a.server.ShutdownHTTP(); err != nil {
		return err
	}

	return a.server.ShutdownGRPC()
}

func (a *application) RunCancelUnpaidOrderWorker(ctx context.Context) error {
	return a.depWorker.cancelUnpaidOrder.Run(ctx)
}
