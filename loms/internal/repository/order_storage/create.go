package orderstorage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/repository/order_storage/sqlc"
)

func (r *repository) Create(ctx context.Context, user int64, items []*OrderItem) (int64, error) {
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

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	return orderID, nil
}
