package orderstorage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
	"route256.ozon.ru/project/loms/internal/repository/order_storage/sqlc"
)

func (r *repository) GetByID(ctx context.Context, orderID int64) (*Order, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, "orderStorage.GetByID")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"GetByID",
		"select",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	queries := sqlc.New(r.dbClient.GetReaderPoolByOrderID(orderID))
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
			metrics.UpdateDatabaseResponseCode(
				RepositoryName,
				"GetByID",
				"select",
				"error",
			)
			return nil, fmt.Errorf("failed to select order %d: %w", orderID, err)
		}
	}

	orderItems := make([]*sqlc.OrderItem, 0, len(items))
	for _, item := range items {
		orderItems = append(orderItems, &item)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"GetByID",
		"select",
		"ok",
	)

	return &Order{
		ID:     order.ID,
		User:   order.User,
		Status: string(order.Status),
		Items:  toOrderItems(orderItems),
	}, nil
}
