package internal

import (
	"context"
	"fmt"
	"time"

	"route256/loms/internal/config"
	"route256/loms/internal/repository"
	"route256/loms/internal/repository/memstore"
	"route256/loms/internal/server"
	"route256/loms/internal/service/order"
	"route256/loms/internal/service/stock"
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
	a.server, err = server.New(
		a.config.Server,
		a.depService.order,
		a.depService.stock,
	)
	if err != nil {
		return err
	}

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
