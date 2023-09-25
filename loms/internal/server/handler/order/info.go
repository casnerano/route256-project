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

func (h *Handler) Info(ctx context.Context, in *pb.InfoRequest) (*pb.InfoResponse, error) {
	if in.GetOrderId() == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
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
