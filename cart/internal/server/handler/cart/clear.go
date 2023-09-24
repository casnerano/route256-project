package cart

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "route256/cart/api/v1"
	cartService "route256/cart/internal/service/cart"
	"time"
)

func (s Handler) Clear(ctx context.Context, in *pb.ClearRequest) (*empty.Empty, error) {
	response := &empty.Empty{}

	if in.GetUser() == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	err := s.service.Clear(sCtx, in.GetUser())
	if err != nil {
		if errors.Is(err, cartService.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	return response, nil
}
