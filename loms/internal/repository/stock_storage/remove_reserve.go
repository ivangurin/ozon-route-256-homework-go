package stockstorage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/repository/stock_storage/sqlc"
)

func (r *repository) RemoveReserve(ctx context.Context, items ReserveItems) error {
	pool := r.dbClient.GetWriterPool()

	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		qtx := sqlc.New(pool).WithTx(tx)

		for _, item := range items {
			stock, err := qtx.GetBySKU(ctx, item.Sku)
			if err != nil {
				return fmt.Errorf("failed to get stock for %d sku: %w", item.Sku, err)
			}

			if stock.TotalCount < int32(item.Quantity) {
				return fmt.Errorf("insufficient stock for product with sku %d", item.Sku)
			}

			if stock.Reserved < int32(item.Quantity) {
				return fmt.Errorf("insufficient reserve for product with sku %d", item.Sku)
			}

			err = qtx.RemoveReserve(ctx, sqlc.RemoveReserveParams{Sku: item.Sku, TotalCount: int32(item.Quantity)})
			if err != nil {
				return fmt.Errorf("failed to remove reserve for sku %d: %w", item.Sku, err)
			}
		}
		
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to remove reserve: %w", err)
	}

	return nil
}
