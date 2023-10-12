package internal

import (
	"context"
	"fmt"
	"route256/cart/internal/config"
	"route256/cart/internal/repository"
	"route256/cart/internal/repository/memstore"
	"route256/cart/internal/repository/sqlstore"
	"route256/cart/internal/server"
	"route256/cart/internal/service/cart"
	"route256/cart/internal/service/client/loms"
	"route256/cart/internal/service/client/pim"
	"route256/cart/pkg/limiter"

	"github.com/jackc/pgx/v5/pgxpool"
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
	pimLimiter    *limiter.Limiter
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

	var cartRepo repository.Cart

	if app.config.Database.DSN != "" {
		var pool *pgxpool.Pool
		pool, err = pgxpool.New(context.TODO(), app.config.Database.DSN)
		if err != nil {
			return nil, err
		}
		cartRepo = sqlstore.NewCartRepository(pool)
	} else {
		cartRepo = memstore.NewCartRepository()
	}

	app.pimLimiter, err = limiter.New(app.config.PIM.RateLimiterAddr)
	if err != nil {
		return nil, err
	}

	app.depRepository = &depRepository{
		cart: cartRepo,
	}

	pimClient, err := pim.NewClient(app.config.PIM.Addr, app.pimLimiter)
	if err != nil {
		return nil, err
	}

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

	err = app.init()
	if err != nil {
		return nil, fmt.Errorf("failed init server: %w", err)
	}

	return &app, nil
}

func (a *application) init() error {
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
		return err
	}

	if err := a.pimLimiter.Close(); err != nil {
		return err
	}

	if err := a.server.ShutdownHTTP(); err != nil {
		return err
	}

	return a.server.ShutdownGRPC()
}
