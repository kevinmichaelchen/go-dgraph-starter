package obs

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// The otel package is still in its early phases,
// so it frequently breaks.
// To keep maintenance costs low, I'm trying to keep
// otel library imports confined to this package.

func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanOption) (context.Context, trace.Span) {
	// Tracer creates a named tracer that implements Tracer interface.
	// If the name is an empty string then provider uses default name.
	tr := otel.Tracer("")

	// Start creates a span.
	return tr.Start(ctx, spanName, opts...)
}

func SetError(span trace.Span, err error) {
	span.SetStatus(codes.Error, err.Error())
}

func WithString(span trace.Span, key string, val string) {
	span.SetAttributes(attribute.String(key, val))
}
