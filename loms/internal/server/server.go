package server

import (
	"context"
	"net/http"
	"time"

	"route256/loms/internal/config"
	"route256/loms/internal/repository/memstore"
	hOrder "route256/loms/internal/server/handler/order"
	hStock "route256/loms/internal/server/handler/stock"
	sOrder "route256/loms/internal/service/order"
	sStock "route256/loms/internal/service/stock"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type handler struct {
	stock *hStock.Handler
	order *hOrder.Handler
}

type server struct {
	config     config.Server
	httpServer *http.Server
}

func New(c config.Server) (*server, error) {
	s := &server{
		config: c,
		httpServer: &http.Server{
			Addr: c.Addr,
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
	repStock := memstore.NewStockRepository()
	repOrder := memstore.NewOrderRepository()

	ss := sStock.New(repStock)
	so := sOrder.New(repOrder, repStock)

	return &handler{
		stock: hStock.NewHandler(ss),
		order: hOrder.NewHandler(so),
	}
}

func (s *server) init() {
	h := s.getHandler()

	router := chi.NewRouter()
	router.Use(
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Recoverer,
	)

	router.Post("/api/stock/info", h.stock.Info)

	router.Post("/api/order/create", h.order.Create)
	router.Post("/api/order/info", h.order.Info)
	router.Post("/api/order/pay", h.order.Pay)
	router.Post("/api/order/cancel", h.order.Cancel)

	s.httpServer.Handler = router
}
