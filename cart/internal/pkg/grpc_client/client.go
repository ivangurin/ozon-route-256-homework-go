package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"route256.ozon.ru/project/cart/internal/pkg/grpc_client/middleware"
)

func GetClientConn(ctx context.Context, host string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(middleware.Tracer),
		// grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		panic(err)
	}
	return conn
}
