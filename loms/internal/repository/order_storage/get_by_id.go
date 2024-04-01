package orderstorage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/repository/order_storage/sqlc"
)

func (r *repository) GetByID(ctx context.Context, orderID int64) (*Order, error) {
	queries := sqlc.New(r.dbClient.GetReaderPool())

	order, err := queries.GetOrderByID(ctx, orderID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrNotFound
		} else {
			return nil, fmt.Errorf("failed to select order %d: %w", orderID, err)
		}
	}

	items, err := queries.GetOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("failed to select order %d: %w", orderID, err)
		}
	}

	return &Order{
		ID:     order.ID,
		User:   order.User,
		Status: string(order.Status),
		Items:  toOrderItems(items),
	}, nil
}

func toOrderItems(items []sqlc.GetOrderItemsByOrderIDRow) OrderItems {
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
