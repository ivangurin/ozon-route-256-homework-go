package orderservice

import (
	"context"

	"route256.ozon.ru/project/loms/internal/model"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

type Service interface {
	Create(ctx context.Context, user int64, items model.OrderItems) (int64, error)
	Info(ctx context.Context, orderID int64) (*model.Order, error)
	Pay(ctx context.Context, orderID int64) error
	Cancel(ctx context.Context, orderID int64) error
}

type service struct {
	stockStorage stockstorage.Repository
	orderStorage orderstorage.Repository
}

func NewService(
	stockStorage stockstorage.Repository,
	orderStorage orderstorage.Repository,
) Service {
	return &service{
		stockStorage: stockStorage,
		orderStorage: orderStorage,
	}
}
