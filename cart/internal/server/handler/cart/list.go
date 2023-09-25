package cart

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "route256/cart/internal/server/proto/cart"
	cartService "route256/cart/internal/service/cart"
	"time"
)

func (s Handler) List(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {
	response := &pb.ListResponse{}

	if in.GetUser() == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	list, err := s.service.List(sCtx, in.GetUser())
	if err != nil {
		if errors.Is(err, cartService.ErrNotFound) {
			return response, nil
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
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
