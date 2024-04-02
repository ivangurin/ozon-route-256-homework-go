package stockservice

import (
	"context"

	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

type Service interface {
	Info(ctx context.Context, sku int64) (uint16, error)
}

type service struct {
	stockStorage stockstorage.Repository
}

func NewService(
	stockStorage stockstorage.Repository,
) Service {
	return &service{
		stockStorage: stockStorage,
	}
}
