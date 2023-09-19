package server

import (
	"context"
	"net/http"
	"route256/loms/internal/config"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type handler struct {
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
	return &handler{}
}

func (s *server) init() {
	//h := s.getHandler()

	router := chi.NewRouter()
	router.Use(
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Recoverer,
	)

	//router.Post("/api/cart/item/add", )

	s.httpServer.Handler = router
}
