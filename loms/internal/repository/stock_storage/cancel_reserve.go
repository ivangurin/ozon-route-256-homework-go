package stockstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
	"route256.ozon.ru/project/loms/internal/repository/stock_storage/sqlc"
)

func (r *repository) CancelReserve(ctx context.Context, items ReserveItems) error {
	ctx, span := tracer.StartSpanFromContext(ctx, "stockRepository:CancelReserve")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"CancelReserve",
		"update",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	pool := r.dbClient.GetWriterPool()
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		qtx := sqlc.New(pool).WithTx(tx)

		for _, item := range items {
			stock, err := qtx.GetBySKU(ctx, item.Sku)
			if err != nil {
				return fmt.Errorf("failed to get stock for %d sku: %w", item.Sku, err)
			}

			if stock.TotalCount < int32(item.Quantity) {
				return fmt.Errorf("insufficient stock for product fro sku %d", item.Sku)
			}

			if stock.Reserved < int32(item.Quantity) {
				return fmt.Errorf("insufficient reserve for product for sku %d", item.Sku)
			}

			err = qtx.CancelReserve(ctx, sqlc.CancelReserveParams{Sku: item.Sku, Reserved: int32(item.Quantity)})
			if err != nil {
				return fmt.Errorf("failed to remove reserve for sku %d: %w", item.Sku, err)
			}

		}

		return nil
	})
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"CancelReserve",
			"update",
			"error",
		)
		return fmt.Errorf("failed to cancel reserve: %w", err)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"CancelReserve",
		"update",
		"ok",
	)

	return nil
}
