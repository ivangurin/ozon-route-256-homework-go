package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func Tracer(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, span := tracer.StartSpanFromContext(ctx, method)
	defer span.End()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-trace-id", tracer.GetTraceID(ctx))
	ctx = metadata.AppendToOutgoingContext(ctx, "x-span-id", tracer.GetSpanID(ctx))
	return invoker(ctx, method, req, reply, cc, opts...)
}
