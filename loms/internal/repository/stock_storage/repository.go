package stockstorage

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

type Repository interface {
	GetBySku(sku int64) (uint16, error)
	Reserve(items ReserveItems) error
	RemoveReserve(items ReserveItems) error
	CancelReserve(items ReserveItems) error
}

type repository struct {
	sync.RWMutex
	ctx   context.Context
	stock Stock
}

func NewRepository(ctx context.Context) Repository {
	r := &repository{
		ctx:   ctx,
		stock: Stock{},
	}

	stockJson, err := os.ReadFile("stock-data.json")
	if err != nil {
		logger.Error("faild to read stocks from file", err)
	} else {
		stockItems := []StockItem{}
		err := json.Unmarshal(stockJson, &stockItems)
		if err != nil {
			logger.Error("faild unmarshal stocks from json", err)
		} else {
			for _, stockItem := range stockItems {
				r.stock[stockItem.Sku] = &stockItem
			}
		}
	}

	return r
}
