package order

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256/loms/internal/model"
	pb "route256/loms/internal/server/proto/order"
	orderService "route256/loms/internal/service/order"
	"time"
)

func (h *Handler) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	if in.GetUser() == 0 && len(in.GetItems()) == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	orderItems := make([]*model.OrderItem, 0, len(in.GetItems()))
	for _, rItem := range in.GetItems() {
		orderItems = append(orderItems, &model.OrderItem{
			SKU:   rItem.GetSku(),
			Count: rItem.GetCount(),
		})
	}

	createdOrder, err := h.service.Create(sCtx, in.GetUser(), orderItems)
	if err != nil {
		if errors.Is(err, orderService.ErrAddReserve) {
			return nil, status.Error(codes.Unknown, err.Error())
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	response := &pb.CreateResponse{
		OrderId: createdOrder.ID,
	}

	return response, nil
}
