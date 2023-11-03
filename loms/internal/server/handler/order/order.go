package order

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"route256/loms/internal/model"
	pb "route256/loms/pkg/proto/order/v1"
)

type Service interface {
	Create(ctx context.Context, userID model.UserID, items []*model.OrderItem) (*model.Order, error)
	GetInfo(ctx context.Context, orderID model.OrderID) (*model.Order, error)
	Payment(ctx context.Context, orderID model.OrderID) error
	Cancel(ctx context.Context, orderID model.OrderID) error
}

type Handler struct {
	pb.UnimplementedOrderServer
	service Service
	logger  *zap.Logger
}

func NewHandler(service Service, logger *zap.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

func (h *Handler) transformStatus(s model.OrderStatus) (pb.Status, error) {
	switch s {
	case model.OrderStatusNew:
		return pb.Status_NEW, nil
	case model.OrderStatusAwaitPayment:
		return pb.Status_AWAIT_PAYMENT, nil
	case model.OrderStatusFailed:
		return pb.Status_FAILED, nil
	case model.OrderStatusPayed:
		return pb.Status_PAYED, nil
	case model.OrderStatusCanceled:
		return pb.Status_CANCELED, nil
	default:
		return 0, errors.New("unknown status")
	}
}
