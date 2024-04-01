package stockstorage

import (
	"context"

	"route256.ozon.ru/project/loms/internal/db"
)

type Repository interface {
	GetBySku(ctx context.Context, sku int64) (uint16, error)
	Reserve(ctx context.Context, items ReserveItems) error
	RemoveReserve(ctx context.Context, items ReserveItems) error
	CancelReserve(ctx context.Context, items ReserveItems) error
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
