package stock

import (
	"context"
	"route256/loms/internal/model"
	pb "route256/loms/pkg/proto/stock/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	GetAvailable(ctx context.Context, sku model.SKU) (uint32, error)
}

type Handler struct {
	pb.UnimplementedStockServer
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Info(ctx context.Context, in *pb.InfoRequest) (*pb.InfoResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	var err error
	response := &pb.InfoResponse{}

	response.Count, err = h.service.GetAvailable(sCtx, in.GetSku())
	if err != nil {
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	return response, nil
}
