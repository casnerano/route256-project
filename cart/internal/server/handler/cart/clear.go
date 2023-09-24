package cart

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "route256/cart/api/v1"
)

func (s Handler) Clear(context.Context, *pb.ClearRequest) (*empty.Empty, error) {
	return nil, nil
}
