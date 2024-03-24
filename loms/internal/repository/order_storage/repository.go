package orderstorage

import (
	"context"
	"sync"

	"route256.ozon.ru/project/loms/internal/db"
)

type Repository interface {
	Create(ctx context.Context, user int64, items []*OrderItem) (int64, error)
	SetStatus(ctx context.Context, orderID int64, status string) error
	GetByID(ctx context.Context, orderID int64) (*Order, error)
}

type repository struct {
	sync.RWMutex
	ctx      context.Context
	dbClient db.Client
	orders   Orders
}

func NewRepository(ctx context.Context, dbClient db.Client) Repository {
	return &repository{
		ctx:      ctx,
		dbClient: dbClient,
		orders:   Orders{},
	}
}
