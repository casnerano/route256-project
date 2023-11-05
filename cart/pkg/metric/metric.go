package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	CounterRequestsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "app_requests_total",
		},
	)

	CounterErrorsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "app_errors_total",
		},
	)

	HistogramResponseTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "app_response_time_seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"code_response"},
	)
)
