package orderstorage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
	"route256.ozon.ru/project/loms/internal/repository/order_storage/sqlc"
)

func (r *repository) GetByIDs(ctx context.Context, orderIDs []int64) ([]*Order, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, "orderStorage.GetByIDs")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"GetByIDs",
		"select",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	shardIDs := map[int64][]int64{}
	for _, orderID := range orderIDs {
		shard := r.dbClient.GetShardByOrderID(orderID)
		shardIDs[shard] = append(shardIDs[shard], orderID)
	}

	res := make([]*Order, 0, len(orderIDs))
	for shardID, orderIDs := range shardIDs {

		pool := r.dbClient.GetReaderPoolByShadID(shardID)
		queries := sqlc.New(pool)
		orders, err := queries.GetOrderByIDs(ctx, orderIDs)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				continue
			} else {
				metrics.UpdateDatabaseResponseCode(
					RepositoryName,
					"GetByIDs",
					"select",
					"error",
				)
				return nil, fmt.Errorf("failed to select order %d: %w", err)
			}
		}

		orderItems, err := queries.GetOrderItemsByOrderIDs(ctx, orderIDs)
		if err != nil {
			if !errors.Is(err, pgx.ErrNoRows) {
				metrics.UpdateDatabaseResponseCode(
					RepositoryName,
					"GetByIDs",
					"select",
					"error",
				)
				return nil, fmt.Errorf("failed to select order items: %w", err)
			}
		}

		orderItemsMap := map[int64][]*sqlc.OrderItem{}
		for _, orderItem := range orderItems {
			orderItemsMap[orderItem.OrderID] = append(orderItemsMap[orderItem.OrderID], &orderItem)
		}

		for _, orders := range orders {
			res = append(res, toOrder(orders, orderItemsMap[orders.ID]))
		}

	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"GetByID",
		"select",
		"ok",
	)

	return res, nil
}

func toOrder(order sqlc.Order, items []*sqlc.OrderItem) *Order {
	return &Order{
		ID:     order.ID,
		User:   order.User,
		Status: string(order.Status),
		Items:  toOrderItems(items),
	}
}

func toOrderItems(items []*sqlc.OrderItem) OrderItems {
	res := make(OrderItems, 0, len(items))
	for _, item := range items {
		res = append(res, &OrderItem{
			ID:       item.ID,
			Sku:      item.Sku,
			Quantity: uint16(item.Quantity),
		})
	}
	return res
}
