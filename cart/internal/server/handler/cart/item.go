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

func (s Handler) ItemAdd(ctx context.Context, in *pb.ItemAddRequest) (*empty.Empty, error) {
	response := &empty.Empty{}

	if in.GetUser() == 0 || in.GetCount() == 0 || in.GetSku() == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	err := s.service.Add(sCtx, in.GetUser(), in.GetSku(), in.GetCount())
	if err != nil {
		if errors.Is(err, cartService.ErrPIMProductNotFound) || errors.Is(err, cartService.ErrPIMLowAvailability) {
			return nil, status.Error(codes.Unknown, err.Error())
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	return response, nil
}

func (s Handler) ItemDelete(ctx context.Context, in *pb.ItemDeleteRequest) (*empty.Empty, error) {
	response := &empty.Empty{}

	if in.GetUser() == 0 || in.GetSku() == 0 {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	err := s.service.Delete(sCtx, in.GetUser(), in.GetSku())
	if err != nil {
		if errors.Is(err, cartService.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}

		return nil, status.Error(codes.Internal, codes.Internal.String())
	}

	return response, nil
}
