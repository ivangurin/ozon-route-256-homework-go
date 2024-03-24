package serviceprovider

import (
	"context"

	orderapi "route256.ozon.ru/project/loms/internal/api/order"
	stockapi "route256.ozon.ru/project/loms/internal/api/stock"
)

type api struct {
	orderAPI *orderapi.API
	stockAPI *stockapi.API
}

func (sp *ServiceProvider) GetOrderAPI(ctx context.Context) *orderapi.API {
	if sp.api.orderAPI == nil {
		sp.api.orderAPI = orderapi.NewAPI(
			sp.GetOrderService(ctx),
		)
	}
	return sp.api.orderAPI
}

func (sp *ServiceProvider) GetStockAPI(ctx context.Context) *stockapi.API {
	if sp.api.stockAPI == nil {
		sp.api.stockAPI = stockapi.NewAPI(
			sp.GetStockService(ctx),
		)
	}
	return sp.api.stockAPI
}
