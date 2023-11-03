package cart

import (
    "context"
    "errors"
    "go.uber.org/zap"
    cartService "route256/cart/internal/service/cart"
    "route256/cart/internal/service/cart/worker_pool"
    pb "route256/cart/pkg/proto/cart/v1"
    "time"

    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func (s Handler) List(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {
    response := &pb.ListResponse{}

    if err := in.ValidateAll(); err != nil {
        return nil, status.Error(codes.InvalidArgument, err.Error())
    }

    sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
    defer cancel()

    list, err := s.service.List(sCtx, worker_pool.New(), in.GetUser())
    if err != nil {
        if errors.Is(err, cartService.ErrNotFound) {
            s.logger.Debug("Cart not found.")

            return response, nil
        }

        s.logger.Warn("Internal error.", zap.Error(err))
        return nil, status.Error(codes.Internal, err.Error())
    }

    for key := range list {
        item := pb.ListItem{
            Sku:   list[key].SKU,
            Count: list[key].Count,
            Name:  list[key].Name,
            Price: list[key].Price,
        }

        response.Items = append(response.Items, &item)
        response.TotalPrice += uint64(item.Price * item.Count)
    }

    return response, nil
}
