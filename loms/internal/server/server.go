package server

import (
	"context"
	"net/http"
	"route256/loms/internal/config"
	handlerOrder "route256/loms/internal/server/handler/order"
	handlerStock "route256/loms/internal/server/handler/stock"
	serviceOrder "route256/loms/internal/service/order"
	serviceStock "route256/loms/internal/service/stock"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type handler struct {
	order *handlerOrder.Handler
	stock *handlerStock.Handler
}

type Server struct {
	config     config.Server
	handler    *handler
	httpServer *http.Server
}

func New(
	c config.Server,
	serviceOrder *serviceOrder.Order,
	serviceStock *serviceStock.Stock,
) (*Server, error) {
	s := &Server{
		config: c,
		handler: &handler{
			order: handlerOrder.NewHandler(serviceOrder),
			stock: handlerStock.NewHandler(serviceStock),
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

	router.Post("/api/stock/info", s.handler.stock.Info)

	router.Post("/api/order/create", s.handler.order.Create)
	router.Post("/api/order/info", s.handler.order.Info)
	router.Post("/api/order/pay", s.handler.order.Pay)
	router.Post("/api/order/cancel", s.handler.order.Cancel)

	s.httpServer.Handler = router
}
