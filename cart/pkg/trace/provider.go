package trace

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func NewTraceProvider(serviceName string) (*sdktrace.TracerProvider, error) {
	extraResource, _ := resource.New(
		context.TODO(),
		resource.WithOS(),
		resource.WithProcess(),
		resource.WithContainer(),
		resource.WithHost(),
		resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)),
	)

	r, err := resource.Merge(
		resource.Default(),
		extraResource,
	)
	if err != nil {
		return nil, err
	}

	exporter, err := NewJaegerExporter("http://jaeger:14268/api/traces")
	if err != nil {
		return nil, fmt.Errorf("initialize exporter: %w", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(r),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp, nil
}
