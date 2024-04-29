package orderservice

import (
	"context"

	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/tracer"
)

func (s *service) GetByIDs(ctx context.Context, orderIDs []int64) ([]*model.Order, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, "orderService.GetByIDs")
	defer span.End()

	orders, err := s.orderStorage.GetByIDs(ctx, orderIDs)
	if err != nil {
		return nil, err
	}

	return ToModelOrders(orders), nil
}
