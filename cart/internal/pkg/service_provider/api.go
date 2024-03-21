package serviceprovider

import (
	cartapi "route256.ozon.ru/project/cart/internal/api/cart_api"
)

type api struct {
	cartAPI cartapi.IAPI
}

func (sp *ServiceProvider) GetCartAPI() cartapi.IAPI {
	if sp.api.cartAPI == nil {
		sp.api.cartAPI = cartapi.NewAPI(
			sp.GetCartService(),
		)
	}
	return sp.api.cartAPI
}
