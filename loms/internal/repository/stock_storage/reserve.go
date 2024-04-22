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

func (r *repository) Reserve(ctx context.Context, items ReserveItems) error {
	ctx, span := tracer.StartSpanFromContext(ctx, "stockRepository:Reserve")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"Reserve",
		"update",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	pool := r.dbClient.GetWriterPoolByShadID(0)
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		qtx := sqlc.New(pool).WithTx(tx)

		for _, item := range items {
			stock, err := qtx.GetBySKU(ctx, item.Sku)
			if err != nil {
				return fmt.Errorf("failed to get stock for %d sku: %w", item.Sku, err)
			}

			if stock.TotalCount-stock.Reserved < int32(item.Quantity) {
				return fmt.Errorf("no free stock for sku %d", item.Sku)
			}

			err = qtx.Reserve(ctx, sqlc.ReserveParams{Sku: item.Sku, Reserved: int32(item.Quantity)})
			if err != nil {
				return fmt.Errorf("failed to reserve stock for sku %d: %w", item.Sku, err)
			}
		}

		return nil
	})
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"Reserve",
			"update",
			"error",
		)
		return fmt.Errorf("failed to reserve stock: %w", err)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"Reserve",
		"update",
		"ok",
	)
	return nil
}
