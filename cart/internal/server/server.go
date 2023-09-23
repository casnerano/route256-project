package server

import (
	"context"
	"net/http"
	"route256/cart/internal/config"
	handlerCart "route256/cart/internal/server/handler/cart"
	serviceCart "route256/cart/internal/service/cart"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type handler struct {
	cart *handlerCart.Handler
}

type Server struct {
	config     config.Server
	handler    *handler
	httpServer *http.Server
}

func New(c config.Server, serviceCart *serviceCart.Cart) (*Server, error) {
	s := &Server{
		config: c,
		handler: &handler{
			cart: handlerCart.NewHandler(serviceCart),
		},
		httpServer: &http.Server{
			Addr: c.Addr,
		},
	}

	s.init()

	return s, nil
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Server) init() {
	router := chi.NewRouter()
	router.Use(
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Recoverer,
	)

	router.Post("/api/cart/item/add", s.handler.cart.Add)
	router.Post("/api/cart/item/delete", s.handler.cart.Delete)

	router.Post("/api/cart/list", s.handler.cart.List)
	router.Post("/api/cart/clear", s.handler.cart.Clear)
	router.Post("/api/cart/checkout", s.handler.cart.Checkout)

	s.httpServer.Handler = router
}
