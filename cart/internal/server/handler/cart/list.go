package cart

import (
	"context"
	pb "route256/cart/api/v1"
)

func (s Handler) List(context.Context, *pb.ListRequest) (*pb.ListResponse, error) {
	return nil, nil
}
