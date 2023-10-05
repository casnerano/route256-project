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

func (h *Handler) Info(ctx context.Context, in *pb.InfoRequest) (*pb.InfoResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	foundOrder, err := h.service.GetInfo(sCtx, in.GetOrderId())
	if err != nil {
		if errors.Is(err, orderService.ErrNotFound) {
			return nil, status.Error(codes.Unknown, err.Error())
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	modelStatus, err := h.transformStatus(foundOrder.Status)
	if err != nil {
		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	response := &pb.InfoResponse{
		Status: modelStatus,
		User:   foundOrder.User,
		Items:  nil,
	}

	for _, rItem := range foundOrder.Items {
		response.Items = append(response.Items, &pb.Item{
			Sku:   rItem.SKU,
			Count: rItem.Count,
		})
	}

	return response, nil
}
