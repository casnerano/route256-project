package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"route256/cart/internal/config"
)

type Server struct {
	config   config.Server
	listener net.Listener
	grpc     *grpc.Server
}

func New(c config.Server) (*Server, error) {
	listener, err := net.Listen("tcp", c.Addr)
	if err != nil {
		return nil, err
	}

	s := &Server{
		config:   c,
		listener: listener,
		grpc:     grpc.NewServer(),
	}

	s.init()

	return s, nil
}

func (s *Server) Modifier(modify func(*grpc.Server)) {
	modify(s.grpc)
}

func (s *Server) Run() error {
	return s.grpc.Serve(s.listener)
}

func (s *Server) Shutdown() error {
	s.grpc.GracefulStop()
	return s.listener.Close()
}

func (s *Server) init() {
	reflection.Register(s.grpc)
}
