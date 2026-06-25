// Example: OTel span processor with secret detection.
//
//	go run github.com/datamaskio/examples/otel
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"

	"github.com/datamaskio/maskotel"
)

func main() {
	// Wrap any SpanProcessor with secret detection.
	exporter := tracetest.NewInMemoryExporter()
	base := trace.NewSimpleSpanProcessor(exporter)
	processor := maskotel.NewProcessor(base,
		maskotel.WithLogger(slog.New(slog.NewTextHandler(os.Stderr, nil))),
	)

	tp := trace.NewTracerProvider(trace.WithSpanProcessor(processor))
	otel.SetTracerProvider(tp)
	tracer := tp.Tracer("example")

	ctx := context.Background()
	_, span := tracer.Start(ctx, "api-request")
	span.SetAttributes(
		attribute.String("user", "alice"),
		attribute.String("token", "ghp_1234567890abcdefghijklmnopqrstuvwxyz"),
	)
	span.End()

	// The processor logged a warning about the GitHub token.
	fmt.Println("Span exported — check stderr for secret detection warnings.")
}
