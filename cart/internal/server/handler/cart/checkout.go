package cart

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	cartService "route256/cart/internal/service/cart"
	pb "route256/cart/pkg/proto/cart/v1"
	"time"
)

func (s Handler) Checkout(ctx context.Context, in *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	response := &pb.CheckoutResponse{}

	if in.GetUser() == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	var err error
	response.OrderId, err = s.service.Checkout(sCtx, in.GetUser())
	if err != nil {
		if errors.Is(err, cartService.ErrEmptyCart) {
			return nil, status.Error(codes.Unknown, err.Error())
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	return response, nil
}
