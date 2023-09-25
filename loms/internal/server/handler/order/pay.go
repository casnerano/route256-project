package order

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "route256/loms/internal/server/proto/order"
	orderService "route256/loms/internal/service/order"
	"time"
)

func (h *Handler) Pay(ctx context.Context, in *pb.PayRequest) (*pb.PayResponse, error) {
	if in.GetOrderId() == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
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
