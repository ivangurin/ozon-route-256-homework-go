package orderstorage

import "route256.ozon.ru/project/loms/internal/model"

func (r *repository) Create(user int64, items []*OrderItem) (int64, error) {
	r.Lock()
	defer r.Unlock()

	order := &Order{
		ID:     r.getNextID(),
		User:   user,
		Status: model.OrederStatusNew,
		Items:  items,
	}

	r.orders[order.ID] = order

	return order.ID, nil
}
