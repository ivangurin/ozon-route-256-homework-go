package stockstorage

import (
	"fmt"

	"route256.ozon.ru/project/loms/internal/model"
)

func (r *repository) Reserve(items ReserveItems) error {
	r.Lock()
	defer r.Unlock()

	for _, item := range items {

		stockItem, exists := r.stock[item.Sku]
		if !exists {
			return model.ErrNotFound
		}

		if (stockItem.TotalCount - stockItem.Reserved) < item.Quantity {
			return fmt.Errorf("no free stock for %d", item.Sku)
		}

	}

	for _, item := range items {
		r.stock[item.Sku].Reserved += item.Quantity
	}

	return nil
}
