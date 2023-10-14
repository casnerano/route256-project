package interceptor

import (
	"context"
	"fmt"
	"route256/cart/pkg/limiter"

	"google.golang.org/grpc"
)

func ClientUnaryRateLimiter(limiter *limiter.Limiter) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if err := limiter.Wait(); err != nil {
			return fmt.Errorf("failed waiting rate limiter: %w", err)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
