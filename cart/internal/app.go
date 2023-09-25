package internal

import (
	"fmt"
	"route256/cart/internal/config"
	"route256/cart/internal/repository"
	"route256/cart/internal/repository/memstore"
	"route256/cart/internal/server"
	handlerCart "route256/cart/internal/server/handler/cart"
	"route256/cart/internal/service/cart"
	"route256/cart/internal/service/client/loms"
	"route256/cart/internal/service/client/pim"
	pb "route256/cart/pkg/proto/cart/v1"
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

	app.depService = &depService{
		pimClient:  pim.NewClient(app.config.PIM.Addr),
		lomsClient: loms.NewClient(app.config.LOMS.Addr),
	}

	app.depService.cart = cart.New(
		app.depRepository.cart,
		app.depService.pimClient,
		app.depService.lomsClient,
	)

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

	pb.RegisterCartServer(a.server.GRPC, handlerCart.NewHandler(a.depService.cart))

	return nil
}

func (a *application) RunServer() error {
	return a.server.Run()
}

func (a *application) ShutdownServer() error {
	return a.server.Shutdown()
}
