package serviceprovider

import (
	"route256.ozon.ru/project/cart/internal/config"
	"route256.ozon.ru/project/cart/internal/pkg/client"
	lomsservice "route256.ozon.ru/project/cart/internal/pkg/client/loms_service"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
)

type clients struct {
	productService productservice.Client
	lomsService    *lomsservice.Client
}

func (sp *ServiceProvider) GetProductService() productservice.Client {
	if sp.clients.productService == nil {
		sp.clients.productService = productservice.NewClient()
	}
	return sp.clients.productService
}

func (sp *ServiceProvider) GetLomsService() *lomsservice.Client {
	if sp.clients.lomsService == nil {
		sp.clients.lomsService = lomsservice.NewClient(
			client.GetClientConn(sp.ctx, config.LomsServiceGrpcHost))
	}
	return sp.clients.lomsService
}
