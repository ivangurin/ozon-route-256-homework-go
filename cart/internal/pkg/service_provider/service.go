package serviceprovider

import (
	"context"

	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

type services struct {
	cartService cartservice.Service
}

func (sp *ServiceProvider) GetCartService(ctx context.Context) cartservice.Service {
	if sp.services.cartService == nil {
		sp.services.cartService = cartservice.NewService(
			sp.GetProductService(ctx),
			sp.GetCartStorage(),
			sp.GetLomsService(ctx),
		)
	}
	return sp.services.cartService
}
