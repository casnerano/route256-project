package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func ClientUnaryRateLimiter() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Printf("method=%v, req=%v, invoker=%v, opts=%v\n", method, req, invoker, opts)
		return nil
	}
}
