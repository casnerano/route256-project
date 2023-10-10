package server

import (
	"context"
	"net"
	"net/http"
	"route256/cart/internal/config"
	"route256/cart/internal/model"
	handlerCart "route256/cart/internal/server/handler/cart"
	"route256/cart/internal/service/cart/worker_pool"
	pb "route256/cart/pkg/proto/cart/v1"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type service interface {
	Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint32) (*model.Item, error)
	Delete(ctx context.Context, userID model.UserID, sku model.SKU) error
	List(ctx context.Context, wp *worker_pool.WorkerPool, userID model.UserID) ([]*model.ItemDetail, error)
	Clear(ctx context.Context, userID model.UserID) error
	Checkout(ctx context.Context, userID model.UserID) (model.OrderID, error)
}

type Server struct {
	config   config.Server
	listener net.Listener
	grpc     *grpc.Server
	http     *http.Server
	service  service
}

func New(c config.Server, service service) (*Server, error) {
	s := &Server{
		config:  c,
		service: service,
	}

	if err := s.initGRPC(); err != nil {
		return nil, err
	}

	if err := s.initHTTP(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Server) RunGRPC() error {
	return s.grpc.Serve(s.listener)
}

func (s *Server) RunHTTP() error {
	return s.http.ListenAndServe()
}

func (s *Server) ShutdownGRPC() error {
	s.grpc.GracefulStop()
	return nil
}

func (s *Server) ShutdownHTTP() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.http.Shutdown(ctx)
}

func (s *Server) initGRPC() error {
	s.grpc = grpc.NewServer()

	listener, err := net.Listen("tcp", s.config.AddrGRPC)
	if err != nil {
		return err
	}

	s.listener = listener

	reflection.Register(s.grpc)
	pb.RegisterCartServer(s.grpc, handlerCart.NewHandler(s.service))

	return nil
}

func (s *Server) initHTTP() error {
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	mux := http.NewServeMux()

	mux.Handle("/", gwMux)
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/v1/openapiv2/cart_service.swagger.json")
	})

	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", http.FileServer(http.Dir("./web/swagger-ui"))))

	s.http = &http.Server{
		Addr:    s.config.AddrHTTP,
		Handler: mux,
	}

	err := pb.RegisterCartHandlerFromEndpoint(
		context.TODO(),
		gwMux,
		s.config.AddrGRPC,
		opts,
	)

	return err
}
