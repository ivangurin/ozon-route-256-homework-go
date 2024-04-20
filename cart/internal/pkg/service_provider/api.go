package serviceprovider

import (
	"context"

	cartapi "route256.ozon.ru/project/cart/internal/api/cart_api"
)

type api struct {
	cartAPI cartapi.API
}

func (sp *ServiceProvider) GetCartAPI(ctx context.Context) cartapi.API {
	if sp.api.cartAPI == nil {
		sp.api.cartAPI = cartapi.NewAPI(
			sp.GetCartService(ctx),
		)
	}
	return sp.api.cartAPI
}
