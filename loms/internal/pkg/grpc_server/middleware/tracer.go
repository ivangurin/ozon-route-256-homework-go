package middleware

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func Tracer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, _ := metadata.FromIncomingContext(ctx)
	traceID := md.Get("x-trace-id")[0]
	fmt.Println("TRACE ID", traceID)

	if traceID != "" {
		traceIDHex, _ := trace.TraceIDFromHex(traceID)
		spanContext := trace.NewSpanContext(trace.SpanContextConfig{
			TraceID: traceIDHex,
		})
		ctx = trace.ContextWithSpanContext(ctx, spanContext)
	}

	ctx, span := tracer.StartSpanFromContext(ctx, info.FullMethod)
	defer span.End()

	return handler(ctx, req)
}
