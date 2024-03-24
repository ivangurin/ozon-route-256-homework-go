package orderstorage

import (
	"context"
	"sync"
)

type Repository interface {
	Create(user int64, items []*OrderItem) (int64, error)
	SetStatus(orderID int64, status string) error
	GetByID(orderID int64) (*Order, error)
}

type repository struct {
	sync.RWMutex
	orders Orders
}

func NewRepository(ctx context.Context) Repository {
	return &repository{
		orders: Orders{},
	}
}
