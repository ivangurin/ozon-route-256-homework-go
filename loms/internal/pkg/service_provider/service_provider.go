package serviceprovider

import (
	"os"
	"syscall"

	"route256.ozon.ru/project/loms/internal/pkg/closer"
)

type ServiceProvider struct {
	closer closer.ICloser

	api          api
	clients      clients
	services     services
	repositories repositories
}

var serviceProvider *ServiceProvider

func GetServiceProvider() *ServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &ServiceProvider{}
	}
	return serviceProvider
}

func (sp *ServiceProvider) GetCloser() closer.ICloser {
	if sp.closer == nil {
		sp.closer = closer.NewCloser(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	}
	return sp.closer
}
