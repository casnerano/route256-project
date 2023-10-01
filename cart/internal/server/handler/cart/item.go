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

func (s Handler) ItemAdd(ctx context.Context, in *pb.ItemAddRequest) (*pb.ItemAddResponse, error) {
	response := &pb.ItemAddResponse{}

	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	_, err := s.service.Add(sCtx, in.GetUser(), in.GetSku(), in.GetCount())
	if err != nil {
		if errors.Is(err, cartService.ErrPIMProductNotFound) || errors.Is(err, cartService.ErrPIMLowAvailability) {
			return nil, status.Error(codes.Unknown, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}

func (s Handler) ItemDelete(ctx context.Context, in *pb.ItemDeleteRequest) (*pb.ItemDeleteResponse, error) {
	response := &pb.ItemDeleteResponse{}

	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	err := s.service.Delete(sCtx, in.GetUser(), in.GetSku())
	if err != nil {
		if errors.Is(err, cartService.ErrNotFound) {
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}
