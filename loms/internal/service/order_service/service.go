package orderservice

import (
	"context"

	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/kafka"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

type Service interface {
	Create(ctx context.Context, user int64, items model.OrderItems) (int64, error)
	Info(ctx context.Context, orderID int64) (*model.Order, error)
	Pay(ctx context.Context, orderID int64) error
	Cancel(ctx context.Context, orderID int64) error
	GetByIDs(ctx context.Context, orderIDs []int64) ([]*model.Order, error)
}

type service struct {
	stockStorage  stockstorage.Repository
	orderStorage  orderstorage.Repository
	kafkaProducer kafka.Producer
}

func NewService(
	stockStorage stockstorage.Repository,
	orderStorage orderstorage.Repository,
	kafkaProducer kafka.Producer,
) Service {
	return &service{
		stockStorage:  stockStorage,
		orderStorage:  orderStorage,
		kafkaProducer: kafkaProducer,
	}
}
