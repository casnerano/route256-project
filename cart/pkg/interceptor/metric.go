package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"route256/cart/pkg/metric"
	"time"
)

func ServerUnaryMetric() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		h, err := handler(ctx, req)
		if err != nil {
			metric.CounterErrorsTotal.Add(1)
		}

		metric.CounterRequestsTotal.Add(1)
		metric.HistogramResponseTime.WithLabelValues(status.Code(err).String()).Observe(time.Since(start).Seconds())

		return h, err
	}
}
