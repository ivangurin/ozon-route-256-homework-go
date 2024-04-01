package stockstorage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/repository/stock_storage/sqlc"
)

func (r *repository) GetBySku(ctx context.Context, sku int64) (uint16, error) {
	pool := r.dbClient.GetReaderPool()
	queries := sqlc.New(pool)
	stock, err := queries.GetBySKU(ctx, sku)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, model.ErrNotFound
		} else {
			return 0, fmt.Errorf("failed to select stock by sku: %w", err)
		}
	}

	return uint16(stock.TotalCount - stock.Reserved), nil
}
