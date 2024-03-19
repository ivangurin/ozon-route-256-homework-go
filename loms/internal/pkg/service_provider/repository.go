package serviceprovider

import (
	"context"

	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

type repositories struct {
	orderStorage orderstorage.Repository
	stockStorage stockstorage.Repository
}

func (sp *ServiceProvider) GetOrderStorage(ctx context.Context) orderstorage.Repository {
	if sp.repositories.orderStorage == nil {
		sp.repositories.orderStorage = orderstorage.NewRepository(ctx)
	}
	return sp.repositories.orderStorage
}

func (sp *ServiceProvider) GetStockStorage(ctx context.Context) stockstorage.Repository {
	if sp.repositories.stockStorage == nil {
		sp.repositories.stockStorage = stockstorage.NewRepository(ctx)
	}
	return sp.repositories.stockStorage
}
