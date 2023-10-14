package order

import (
	"context"
	"errors"
	orderService "route256/loms/internal/service/order"
	pb "route256/loms/pkg/proto/order/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Pay(ctx context.Context, in *pb.PayRequest) (*pb.PayResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	err := h.service.Payment(sCtx, in.GetOrderId())
	if err != nil {
		if errors.Is(err, orderService.ErrNotFound) || errors.Is(err, orderService.ErrShipReserve) {
			return nil, status.Error(codes.Unknown, err.Error())
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	response := &pb.PayResponse{}

	return response, nil
}
