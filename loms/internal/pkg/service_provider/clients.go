package serviceprovider

import (
	"context"

	"route256.ozon.ru/project/loms/internal/config"
	cartservice "route256.ozon.ru/project/loms/internal/pkg/client/cart_service"
)

type clients struct {
	cartService cartservice.Client
}

func (sp *ServiceProvider) GetCartService(ctx context.Context) cartservice.Client {
	if sp.clients.cartService == nil {
		sp.clients.cartService = cartservice.NewClient(ctx, config.CartServiceHost)
	}
	return sp.clients.cartService
}
