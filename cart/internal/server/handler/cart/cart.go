package cart

import (
	"context"
	"go.uber.org/zap"
	"route256/cart/internal/model"
	"route256/cart/internal/service/cart/worker_pool"
	pb "route256/cart/pkg/proto/cart/v1"
)

type Service interface {
	Add(ctx context.Context, userID model.UserID, sku model.SKU, count uint32) (*model.Item, error)
	Delete(ctx context.Context, userID model.UserID, sku model.SKU) error
	List(ctx context.Context, wp worker_pool.WorkerPool, userID model.UserID) ([]*model.ItemDetail, error)
	Clear(ctx context.Context, userID model.UserID) error
	Checkout(ctx context.Context, userID model.UserID) (model.OrderID, error)
}

type Handler struct {
	pb.UnimplementedCartServer
	service Service
	logger  *zap.Logger
}

func NewHandler(service Service, logger *zap.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}
