package orderservice

import (
	"context"

	"route256.ozon.ru/project/loms/internal/model"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

type Service interface {
	Create(user int64, items model.OrderItems) (int64, error)
	Info(orderID int64) (*model.Order, error)
	Pay(orderID int64) error
	Cancel(orderID int64) error
}

type service struct {
	stockStorage stockstorage.Repository
	orderStorage orderstorage.Repository
}

func NewService(
	ctx context.Context,
	stockStorage stockstorage.Repository,
	orderStorage orderstorage.Repository,
) Service {
	return &service{
		stockStorage: stockStorage,
		orderStorage: orderStorage,
	}
}
