package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"route256/loms/internal/config"
)

type Server struct {
	config   config.Server
	listener net.Listener
	GRPC     *grpc.Server
}

func New(c config.Server) (*Server, error) {
	listener, err := net.Listen("tcp", c.Addr)
	if err != nil {
		return nil, err
	}

	s := &Server{
		config:   c,
		listener: listener,
		GRPC:     grpc.NewServer(),
	}

	s.init()

	return s, nil
}

func (s *Server) Run() error {
	return s.GRPC.Serve(s.listener)
}

func (s *Server) Shutdown() error {
	s.GRPC.GracefulStop()
	return s.listener.Close()
}

func (s *Server) init() {
	reflection.Register(s.GRPC)
}
