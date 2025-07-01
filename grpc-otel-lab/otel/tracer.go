package otel

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/exporters/jaeger"
)

// InitTracer initialises OTEL and returns TracerProvider + ctx.
func InitTracer(service string) (*trace.TracerProvider, context.Context) {
	exp, _ := jaeger.New(jaeger.WithCollectorEndpoint())
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
		)),
	)
	log.Printf("OpenTelemetry initialised for %s", service)
	return tp, context.Background()
}
