package orderservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func (s *service) Create(ctx context.Context, user int64, items model.OrderItems) (int64, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, "orderService.Create")
	defer span.End()

	orderID, err := s.orderStorage.Create(ctx, user, ToOrderStorageItems(items))
	if err != nil {
		logger.Errorf(ctx, "failed to create order: %v", err)
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	metrics.UpdateOrdersCreated()

	reserved := false
	reserveErr := s.stockStorage.Reserve(ctx, ToStockItems(items))
	if reserveErr != nil {
		logger.Errorf(ctx, "failed to reserve quantity for items: %v", reserveErr)
	} else {
		reserved = true
	}

	var status string
	if reserved {
		status = model.OrderStatusAwaitingPayment
	} else {
		status = model.OrderStatusFailed
	}

	err = s.orderStorage.SetStatus(ctx, orderID, status)
	if err != nil {
		logger.Errorf(ctx, "failed to change status: %w", err)
		return 0, fmt.Errorf("failed to change status: %w", err)
	}

	metrics.UpdateOrderStatusChanged(model.OrderStatusNew, status)

	if !reserved {
		return 0, fmt.Errorf("failed to reserve quantity for items, %w", reserveErr)
	}

	return orderID, nil
}
