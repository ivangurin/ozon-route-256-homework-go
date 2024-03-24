package stockstorage

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"route256.ozon.ru/project/loms/internal/db"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

type Repository interface {
	GetBySku(ctx context.Context, sku int64) (uint16, error)
	Reserve(ctx context.Context, items ReserveItems) error
	RemoveReserve(ctx context.Context, items ReserveItems) error
	CancelReserve(ctx context.Context, items ReserveItems) error
}

type repository struct {
	sync.RWMutex
	ctx      context.Context
	dbClient db.Client
	stock    Stock
}

func NewRepository(ctx context.Context, dbClient db.Client) Repository {
	r := &repository{
		ctx:      ctx,
		dbClient: dbClient,
		stock:    Stock{},
	}

	stockJson, err := os.ReadFile("stock-data.json")
	if err != nil {
		logger.Errorf(ctx, "failed to read stocks from file: %w", err)
	} else {
		stockItems := []StockItem{}
		err := json.Unmarshal(stockJson, &stockItems)
		if err != nil {
			logger.Errorf(ctx, "failed unmarshal stocks from json :%w ", err)
		} else {
			for _, stockItem := range stockItems {
				r.stock[stockItem.Sku] = &stockItem
			}
		}
	}

	return r
}
