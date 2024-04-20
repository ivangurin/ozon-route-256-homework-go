package serviceprovider

import (
	"context"

	"route256.ozon.ru/project/cart/internal/config"
	lomsservice "route256.ozon.ru/project/cart/internal/pkg/client/loms_service"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	client "route256.ozon.ru/project/cart/internal/pkg/grpc_client"
)

type clients struct {
	productService productservice.Client
	lomsService    *lomsservice.Client
}

func (sp *ServiceProvider) GetProductService(ctx context.Context) productservice.Client {
	if sp.clients.productService == nil {
		sp.clients.productService = productservice.NewClient(
			sp.GetRedisClient(ctx),
		)
	}
	return sp.clients.productService
}

func (sp *ServiceProvider) GetLomsService(ctx context.Context) *lomsservice.Client {
	if sp.clients.lomsService == nil {
		sp.clients.lomsService = lomsservice.NewClient(
			client.GetClientConn(ctx, config.LomsServiceGrpcHost))
	}
	return sp.clients.lomsService
}
