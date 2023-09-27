package internal

import (
	"fmt"
	"route256/cart/internal/config"
	"route256/cart/internal/repository"
	"route256/cart/internal/repository/memstore"
	"route256/cart/internal/server"
	"route256/cart/internal/service/cart"
	"route256/cart/internal/service/client/loms"
	"route256/cart/internal/service/client/pim"
)

type depRepository struct {
	cart repository.Cart
}

type depService struct {
	cart       *cart.Cart
	pimClient  *pim.Client
	lomsClient *loms.Client
}

type application struct {
	config        *config.Config
	server        *server.Server
	depRepository *depRepository
	depService    *depService
}

func NewApp() (*application, error) {
	var err error
	app := application{}

	app.config, err = config.New()
	if err != nil {
		return nil, err
	}

	app.depRepository = &depRepository{
		cart: memstore.NewCartRepository(),
	}

	pimClient := pim.NewClient(app.config.PIM.Addr)

	lomsClient, err := loms.NewClient(app.config.LOMS.Addr)
	if err != nil {
		return nil, err
	}

	app.depService = &depService{
		pimClient:  pimClient,
		lomsClient: lomsClient,
	}

	app.depService.cart = cart.New(
		app.depRepository.cart,
		app.depService.pimClient,
		app.depService.lomsClient,
	)

	err = app.initGRPCServer()
	if err != nil {
		return nil, fmt.Errorf("failed init server: %w", err)
	}

	return &app, nil
}

func (a *application) initGRPCServer() error {
	var err error
	a.server, err = server.New(a.config.Server, a.depService.cart)

	return err
}

func (a *application) RunGRPCServer() error {
	return a.server.RunGRPC()
}

func (a *application) RunHTTPServer() error {
	return a.server.RunHTTP()
}

func (a *application) Shutdown() error {
	if err := a.depService.lomsClient.Close(); err != nil {
		return nil
	}

	if err := a.server.ShutdownHTTP(); err != nil {
		return err
	}

	return a.server.ShutdownGRPC()
}
