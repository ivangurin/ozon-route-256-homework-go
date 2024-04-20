package serviceprovider

import (
	"context"
	"os"
	"syscall"

	"route256.ozon.ru/project/cart/internal/pkg/closer"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
	"route256.ozon.ru/project/cart/internal/pkg/redis"
)

type ServiceProvider struct {
	ctx          context.Context
	closer       closer.Closer
	redis        redis.Client
	api          api
	clients      clients
	repositories repositories
	services     services
}

var serviceProvider *ServiceProvider

func GetServiceProvider(ctx context.Context) *ServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &ServiceProvider{
			ctx: ctx,
		}
	}
	return serviceProvider
}

func (sp *ServiceProvider) GetCloser() closer.Closer {
	if sp.closer == nil {
		sp.closer = closer.NewCloser(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	}
	return sp.closer
}

func (sp *ServiceProvider) GetRedisClient(ctx context.Context) redis.Client {
	var err error
	if sp.redis == nil {
		sp.redis, err = redis.NewClient()
		if err != nil {
			logger.Panicf(ctx, "failed to create redis client: %v", err)
		}
	}
	return sp.redis
}
