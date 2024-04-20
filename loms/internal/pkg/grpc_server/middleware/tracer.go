package middleware

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func Tracer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)

	var traceID string
	if traceIDs, exists := md["x-trace-id"]; exists {
		traceID = traceIDs[0]
	}

	var spanID string
	if spanIDs, exists := md["x-span-id"]; exists {
		spanID = spanIDs[0]
	}

	if traceID != "" {
		var err error
		spanContext := trace.SpanContextConfig{
			TraceFlags: trace.FlagsSampled,
			Remote:     true,
		}
		spanContext.TraceID, err = trace.TraceIDFromHex(traceID)
		if err != nil {
			return nil, err
		}
		spanContext.SpanID, err = trace.SpanIDFromHex(spanID)
		if err != nil {
			return nil, err
		}
		ctx = trace.ContextWithSpanContext(ctx,
			trace.NewSpanContext(spanContext))
	}

	var span trace.Span
	ctx, span = tracer.StartSpanFromContext(ctx, info.FullMethod,
		trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	return handler(ctx, req)
}
