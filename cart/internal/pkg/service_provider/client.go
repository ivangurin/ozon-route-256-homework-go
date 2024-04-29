package serviceprovider

import (
	"context"

	"route256.ozon.ru/project/cart/internal/config"
	"route256.ozon.ru/project/cart/internal/pkg/cache"
	"route256.ozon.ru/project/cart/internal/pkg/cache/redis"
	lomsservice "route256.ozon.ru/project/cart/internal/pkg/client/loms_service"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	client "route256.ozon.ru/project/cart/internal/pkg/grpc_client"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

type clients struct {
	productService productservice.Client
	lomsService    *lomsservice.Client
	redisClient    cache.Cache
}

func (sp *ServiceProvider) GetRedisClient(ctx context.Context) cache.Cache {
	var err error
	if sp.clients.redisClient == nil {
		sp.clients.redisClient, err = redis.NewCache(config.RedisUrl)
		if err != nil {
			logger.Panicf(ctx, "failed to create redis client: %v", err)
		}
	}
	return sp.clients.redisClient
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
			client.GetClientConn(
				sp.ctx,
				lomsservice.ServiceName,
				config.LomsServiceGrpcHost))
	}
	return sp.clients.lomsService
}
