package orderservice

import (
	"context"

	"route256.ozon.ru/project/loms/internal/model"
)

func (s *service) Pay(ctx context.Context, orderID int64) error {
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

	return nil
}