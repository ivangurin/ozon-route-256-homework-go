package serviceprovider

import (
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

type services struct {
	cartService cartservice.Service
}

func (sp *ServiceProvider) GetCartService() cartservice.Service {
	if sp.services.cartService == nil {
		sp.services.cartService = cartservice.NewService(
			sp.GetProductService(),
			sp.GetCartStorage(),
			sp.GetLomsService(),
		)
	}
	return sp.services.cartService
}
