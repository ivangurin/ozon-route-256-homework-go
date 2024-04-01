package orderstorage

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/loms/internal/repository/order_storage/sqlc"
)

func (r *repository) SetStatus(ctx context.Context, orderID int64, status string) error {
	queries := sqlc.New(r.dbClient.GetWriterPool())

	err := queries.UpdateStatusByOrderID(ctx, sqlc.UpdateStatusByOrderIDParams{ID: orderID, Status: sqlc.OrderStatusType(status)})
	if err != nil {
		return fmt.Errorf("failed to update order status for %d: %w", orderID, err)
	}

	return nil
}
