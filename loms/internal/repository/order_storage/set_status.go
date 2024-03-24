package orderstorage

import "route256.ozon.ru/project/loms/internal/model"

func (r *repository) SetStatus(orderID int64, status string) error {
	r.Lock()
	defer r.Unlock()

	order, exists := r.orders[orderID]
	if !exists {
		return model.ErrNotFound
	}

	order.Status = status

	return nil
}
