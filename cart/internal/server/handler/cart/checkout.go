package cart

import (
	"context"
	pb "route256/cart/api/v1"
)

func (s Handler) Checkout(context.Context, *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	return nil, nil
}
