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

func (h *Handler) Cancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelResponse, error) {
	if in.GetOrderId() == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	err := h.service.Cancel(sCtx, in.GetOrderId())
	if err != nil {
		if errors.Is(err, orderService.ErrNotFound) ||
			errors.Is(err, orderService.ErrCancelPaidOrder) ||
			errors.Is(err, orderService.ErrCancelReserve) {

			return nil, status.Error(codes.Unknown, err.Error())
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	response := &pb.CancelResponse{}

	return response, nil
}
