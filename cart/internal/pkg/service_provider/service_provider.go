package serviceprovider

import (
	"context"
	"os"
	"syscall"

	"route256.ozon.ru/project/cart/internal/pkg/closer"
)

type ServiceProvider struct {
	ctx          context.Context
	closer       closer.ICloser
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

func (sp *ServiceProvider) GetCloser() closer.ICloser {
	if sp.closer == nil {
		sp.closer = closer.NewCloser(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	}
	return sp.closer
}
