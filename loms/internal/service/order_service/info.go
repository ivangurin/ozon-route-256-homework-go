package orderservice

import (
	"context"

	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func (s *service) Info(ctx context.Context, orderID int64) (*model.Order, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, "orderService.Info")
	defer span.End()

	order, err := s.orderStorage.GetByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return ToModelOrder(order), nil
}
