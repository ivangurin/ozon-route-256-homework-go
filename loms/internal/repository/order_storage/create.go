package orderstorage

import (
	"context"

	"route256.ozon.ru/project/loms/internal/model"
)

func (r *repository) Create(ctx context.Context, user int64, items []*OrderItem) (int64, error) {
	r.Lock()
	defer r.Unlock()

	order := &Order{
		ID:     r.getNextID(),
		User:   user,
		Status: model.OrderStatusNew,
		Items:  items,
	}

	r.orders[order.ID] = order

	return order.ID, nil
}
