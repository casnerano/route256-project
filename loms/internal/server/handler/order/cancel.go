package order

import (
	"context"
	"errors"
	"go.uber.org/zap"
	orderService "route256/loms/internal/service/order"
	pb "route256/loms/pkg/proto/order/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) Cancel(ctx context.Context, in *pb.CancelRequest) (*pb.CancelResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
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

		h.logger.Error("Internal error.", zap.Error(err))
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	response := &pb.CancelResponse{}

	return response, nil
}
