package server

import (
	"context"
	"net/http"
	"route256/cart/internal/config"
	"route256/cart/internal/repository/memstore"
	hCart "route256/cart/internal/server/handler/cart"
	sCart "route256/cart/internal/service/cart"
	"route256/cart/internal/service/client/loms"
	"route256/cart/internal/service/client/pim"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type handler struct {
	cart *hCart.Handler
}

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

func (s *server) getHandler() *handler {
	rep := memstore.NewCartRepository()

	pimClient := pim.NewClient(s.config.PIM.Addr)
	lomsClient := loms.NewClient(s.config.LOMS.Addr)

	sc := sCart.New(rep, pimClient, lomsClient)

	return &handler{
		cart: hCart.NewHandler(sc),
	}
}

func (s *server) init() {
	h := s.getHandler()

	router := chi.NewRouter()
	router.Use(
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Recoverer,
	)

	router.Post("/api/cart/item/add", h.cart.Add)
	router.Post("/api/cart/item/delete", h.cart.Delete)

	router.Post("/api/cart/list", h.cart.List)
	router.Post("/api/cart/clear", h.cart.Clear)
	router.Post("/api/cart/checkout", h.cart.Checkout)

	s.httpServer.Handler = router
}
