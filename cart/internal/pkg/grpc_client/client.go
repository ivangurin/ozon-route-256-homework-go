package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"route256.ozon.ru/project/cart/internal/pkg/grpc_client/middleware"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func GetClientConn(ctx context.Context, serviceName string, host string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(middleware.NewMetricInterceptor(serviceName)),
		grpc.WithUnaryInterceptor(middleware.Tracer),
		// grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		logger.Panicf(ctx, "failed to connect to server: %v", err)
	}
	return conn
}
