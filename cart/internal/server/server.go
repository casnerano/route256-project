package server

import (
	"context"
	"net/http"
	"route256/cart/internal/config"
	"route256/cart/internal/repository/memstore"
	cartHandler "route256/cart/internal/server/handler/cart"
	"route256/cart/internal/service/cart"
	"route256/cart/internal/service/client/loms"
	"route256/cart/internal/service/client/pim"
	"time"
)

type server struct {
	config     *config.Config
	httpServer *http.Server
}

func New() (*server, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}

	s := &server{
		config: c,
		httpServer: &http.Server{
			Addr: c.Server.Addr,
		},
	}

	s.init()

	return s, nil
}

func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *server) init() {
	// Init dependencies
	rep := memstore.NewCartRepository()

	pimClient := pim.NewClient(s.config.PIM.Addr)
	lomsClient := loms.NewClient(s.config.LOMS.Addr)

	cartService := cart.New(rep, pimClient, lomsClient)

	itemHandler := cartHandler.NewItemHandler(cartService)
	listHandler := cartHandler.NewListHandler(cartService)
	clearHandler := cartHandler.NewClearHandler(cartService)
	checkoutHandler := cartHandler.NewCheckoutHandler()

	// Init routes
	mux := http.NewServeMux()

	mux.HandleFunc("/api/cart/item/add", itemHandler.Add)
	mux.HandleFunc("/api/cart/item/delete", itemHandler.Delete)

	mux.HandleFunc("/api/cart/list", listHandler.List)
	mux.HandleFunc("/api/cart/clear", clearHandler.Clear)
	mux.HandleFunc("/api/cart/checkout", checkoutHandler.Checkout)

	s.httpServer.Handler = mux
}
