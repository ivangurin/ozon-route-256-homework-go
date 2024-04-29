package orderstorage

import (
	"context"

	"route256.ozon.ru/project/loms/internal/db"
)

const (
	RepositoryName = "OrderStorage"
)

type Repository interface {
	Create(ctx context.Context, user int64, items []*OrderItem) (int64, error)
	SetStatus(ctx context.Context, orderID int64, status string) error
	GetByID(ctx context.Context, orderID int64) (*Order, error)
	GetByIDs(ctx context.Context, orderID []int64) ([]*Order, error)
}

type repository struct {
	ctx      context.Context
	dbClient db.Client
}

func NewRepository(ctx context.Context, dbClient db.Client) Repository {
	return &repository{
		ctx:      ctx,
		dbClient: dbClient,
	}
}
