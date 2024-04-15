package stockstorage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/repository/stock_storage/sqlc"
)

func (r *repository) GetBySku(ctx context.Context, sku int64) (uint16, error) {

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"GetBySku",
		"select",
	)

	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	pool := r.dbClient.GetReaderPool()
	queries := sqlc.New(pool)
	stock, err := queries.GetBySKU(ctx, sku)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			metrics.UpdateDatabaseResponseCode(
				RepositoryName,
				"GetBySku",
				"select",
				"not_found",
			)
			return 0, model.ErrNotFound
		} else {
			metrics.UpdateDatabaseResponseCode(
				RepositoryName,
				"GetBySku",
				"select",
				"error",
			)
			return 0, fmt.Errorf("failed to select stock by sku: %w", err)
		}
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"GetBySku",
		"select",
		"ok",
	)

	return uint16(stock.TotalCount - stock.Reserved), nil
}
