package orderstorage

import (
	"fmt"

	"route256.ozon.ru/project/loms/internal/model"
)

func (r *repository) GetByID(orderID int64) (*Order, error) {
	r.RLock()
	defer r.RUnlock()

	order, exists := r.orders[orderID]
	if !exists {
		return nil, fmt.Errorf("order %d not found: %w", orderID, model.ErrNotFound)
	}

	return order, nil
}
