package server

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"net"
	"net/http"
	"route256/loms/internal/config"
	"route256/loms/internal/model"
	orderHandler "route256/loms/internal/server/handler/order"
	stockHandler "route256/loms/internal/server/handler/stock"
	pbOrder "route256/loms/pkg/proto/order/v1"
	pbStock "route256/loms/pkg/proto/stock/v1"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type orderService interface {
	Create(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error)
	GetInfo(ctx context.Context, orderID model.OrderID) (*model.Order, error)
	Payment(ctx context.Context, orderID model.OrderID) error
	Cancel(ctx context.Context, orderID model.OrderID) error
	CancelUnpaidWithDuration(ctx context.Context, duration time.Duration) error
}

type stockService interface {
	GetAvailable(ctx context.Context, sku model.SKU) (uint32, error)
	AddReserve(ctx context.Context, sku model.SKU, count uint32) error
	CancelReserve(ctx context.Context, sku model.SKU, count uint32) error
	ShipReserve(ctx context.Context, sku model.SKU, count uint32) error
}

type Server struct {
	config       config.Server
	listener     net.Listener
	grpc         *grpc.Server
	http         *http.Server
	orderService orderService
	stockService stockService
	logger       *zap.Logger
}

func New(c config.Server, order orderService, stock stockService, logger *zap.Logger) (*Server, error) {
	s := &Server{
		config:       c,
		orderService: order,
		stockService: stock,
		logger:       logger,
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
	s.grpc = grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	listener, err := net.Listen("tcp", s.config.AddrGRPC)
	if err != nil {
		return err
	}

	s.listener = listener

	reflection.Register(s.grpc)
	pbOrder.RegisterOrderServer(s.grpc, orderHandler.NewHandler(s.orderService, s.logger))
	pbStock.RegisterStockServer(s.grpc, stockHandler.NewHandler(s.stockService, s.logger))

	return nil
}

func (s *Server) initHTTP() error {
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	mux := http.NewServeMux()

	mux.Handle("/", gwMux)
	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./api/v1/openapiv2/loms.swagger.json")
	})

	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", http.FileServer(http.Dir("./web/swagger-ui"))))

	s.http = &http.Server{
		Addr:    s.config.AddrHTTP,
		Handler: mux,
	}

	err := pbOrder.RegisterOrderHandlerFromEndpoint(context.TODO(), gwMux, s.config.AddrGRPC, opts)
	if err != nil {
		return err
	}

	err = pbStock.RegisterStockHandlerFromEndpoint(context.TODO(), gwMux, s.config.AddrGRPC, opts)

	return err
}
