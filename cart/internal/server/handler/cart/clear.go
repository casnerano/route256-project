package cart

import (
	"context"
	"errors"
	cartService "route256/cart/internal/service/cart"
	pb "route256/cart/pkg/proto/cart/v1"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s Handler) Clear(ctx context.Context, in *pb.ClearRequest) (*pb.ClearResponse, error) {
	response := &pb.ClearResponse{}

	if err := in.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	sCtx, cancel := context.WithTimeout(ctx, 300*time.Second)
	defer cancel()

	err := s.service.Clear(sCtx, in.GetUser())
	if err != nil {
		if errors.Is(err, cartService.ErrNotFound) {
			s.logger.Debug("Cart not found.")
			return nil, status.Error(codes.NotFound, codes.NotFound.String())
		}

		s.logger.Warn("Internal error.", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response, nil
}
