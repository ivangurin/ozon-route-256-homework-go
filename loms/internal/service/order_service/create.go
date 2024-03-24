package orderservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func (s *service) Create(ctx context.Context, user int64, items model.OrderItems) (int64, error) {

	orderID, err := s.orderStorage.Create(ctx, user, ToOrderStorageItems(items))
	if err != nil {
		logger.Errorf(ctx, "failed to create order: %w", err)
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	reserved := false
	reserveErr := s.stockStorage.Reserve(ctx, ToStockItems(items))
	if reserveErr != nil {
		logger.Errorf(ctx, "failed to reserve quantity for items: %w", reserveErr)
	} else {
		reserved = true
	}

	if reserved {
		err = s.orderStorage.SetStatus(ctx, orderID, model.OrderStatusAwaitingPayment)
	} else {
		err = s.orderStorage.SetStatus(ctx, orderID, model.OrderStatusFailed)
	}
	if err != nil {
		logger.Errorf(ctx, "failed to change status: %w", err)
		return 0, fmt.Errorf("failed to change status: %w", err)
	}

	if !reserved {
		return 0, fmt.Errorf("failed to reserve quantity for items, %w", reserveErr)
	}

	return orderID, nil
}
