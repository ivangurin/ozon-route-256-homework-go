package orderstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
	"route256.ozon.ru/project/loms/internal/repository/order_storage/sqlc"
)

func (r *repository) Create(ctx context.Context, user int64, items []*OrderItem) (int64, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, "orderStorage.Create")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"Create",
		"insert",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	pool := r.dbClient.GetWriterPool()
	var orderID int64
	err := pool.BeginFunc(ctx, func(tx pgx.Tx) error {
		qtx := sqlc.New(pool).WithTx(tx)

		var err error
		orderID, err = qtx.CreateOrder(ctx, sqlc.CreateOrderParams{User: user, Status: model.OrderStatusNew})
		if err != nil {
			return fmt.Errorf("failed add row to order table: %w", err)
		}

		for _, item := range items {
			err = qtx.AddOrderItem(ctx, sqlc.AddOrderItemParams{OrderID: orderID, Sku: item.Sku, Quantity: int32(item.Quantity)})
			if err != nil {
				return fmt.Errorf("failed to add row to order_item table for sku %d: %w", item.Sku, err)
			}
		}

		err = r.insertOutboxOrderStatusChanged(ctx, tx, orderID, model.OrderStatusNew)
		if err != nil {
			return fmt.Errorf("failed to insert to kafka outbox record: %w", err)
		}

		return nil
	})
	if err != nil {
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"Create",
			"insert",
			"error",
		)
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"Create",
		"insert",
		"ok",
	)

	return orderID, nil
}
