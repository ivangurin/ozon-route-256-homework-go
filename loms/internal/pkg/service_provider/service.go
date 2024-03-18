package serviceprovider

import (
	"context"

	orderservice "route256.ozon.ru/project/loms/internal/service/order_service"
	stockservice "route256.ozon.ru/project/loms/internal/service/stock_service"
)

type services struct {
	orderService orderservice.Service
	stockService stockservice.Service
}

func (sp *ServiceProvider) GetOrderService(ctx context.Context) orderservice.Service {
	if sp.services.orderService == nil {
		sp.services.orderService = orderservice.NewService(ctx)
	}
	return sp.services.orderService
}

func (sp *ServiceProvider) GetStockService(ctx context.Context) stockservice.Service {
	if sp.services.stockService == nil {
		sp.services.stockService = stockservice.NewService(ctx)
	}
	return sp.services.stockService
}
