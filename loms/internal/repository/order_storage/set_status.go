package orderstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/repository/order_storage/sqlc"
)

func (r *repository) SetStatus(ctx context.Context, orderID int64, status string) error {
	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"SetStatus",
		"update",
	)

	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	pool := r.dbClient.GetWriterPool()
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		qtx := sqlc.New(pool).WithTx(tx)

		var err error
		err = qtx.UpdateStatusByOrderID(ctx, sqlc.UpdateStatusByOrderIDParams{ID: orderID, Status: sqlc.OrderStatusType(status)})
		if err != nil {
			return fmt.Errorf("failed to update order status for %d: %w", orderID, err)
		}

		err = r.insertOutboxOrderStatusChanged(ctx, tx, orderID, status)
		if err != nil {
			return fmt.Errorf("failed to insert to kafka outbox record: %w", err)
		}

		return nil
	})
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"SetStatus",
			"update",
			"error",
		)
		return fmt.Errorf("failed to set status for %d: %w", orderID, err)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"SetStatus",
		"update",
		"ok",
	)

	return nil
}
