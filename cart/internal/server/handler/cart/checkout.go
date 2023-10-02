package cart

import (
	"context"
	"errors"
	cartService "route256/cart/internal/service/cart"
	pb "route256/cart/pkg/proto/cart/v1"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Handler) Checkout(ctx context.Context, in *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	response := &pb.CheckoutResponse{}

	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	var err error
	response.OrderId, err = s.service.Checkout(sCtx, in.GetUser())
	if err != nil {
		if errors.Is(err, cartService.ErrEmptyCart) {
			return nil, status.Error(codes.Unknown, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}
