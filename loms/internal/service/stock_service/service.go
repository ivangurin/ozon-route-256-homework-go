package stockservice

import (
	"context"

	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

type Service interface {
	Info(sku int64) (uint16, error)
}

type service struct {
	stockStorage stockstorage.Repository
}

func NewService(
	ctx context.Context,
	stockStorage stockstorage.Repository,
) Service {
	return &service{
		stockStorage: stockStorage,
	}
}
