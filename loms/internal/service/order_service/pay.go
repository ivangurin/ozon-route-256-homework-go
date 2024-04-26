package orderservice

import (
	"context"

	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func (s *service) Pay(ctx context.Context, orderID int64) error {
	ctx, span := tracer.StartSpanFromContext(ctx, "orderService.Pay")
	defer span.End()

	orderStorage, err := s.orderStorage.GetByID(ctx, orderID)
	if err != nil {
		return err
	}

	order := ToModelOrder(orderStorage)

	err = s.stockStorage.RemoveReserve(ctx, ToStockItems(order.Items))
	if err != nil {
		return err
	}

	err = s.orderStorage.SetStatus(ctx, order.ID, model.OrderStatusPayed)
	if err != nil {
		return err
	}

	metrics.UpdateOrderStatusChanged(orderStorage.Status, model.OrderStatusPayed)

	return nil
}
