package cart

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "route256/cart/api/v1"
)

func (s Handler) ItemAdd(context.Context, *pb.ItemAddRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s Handler) ItemDelete(context.Context, *pb.ItemDeleteRequest) (*empty.Empty, error) {
	return nil, nil
}
