package middleware

import (
	"context"

	"google.golang.org/grpc"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func Tracer(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	ctx, span := tracer.StartSpanFromContext(ctx, info.FullMethod)
	defer span.End()

	return handler(ctx, req)
}
