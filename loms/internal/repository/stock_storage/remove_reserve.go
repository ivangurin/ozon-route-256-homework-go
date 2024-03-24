package stockstorage

import (
	"context"
	"fmt"
)

func (r *repository) RemoveReserve(ctx context.Context, items ReserveItems) error {
	r.Lock()
	defer r.Unlock()

	for _, item := range items {
		stockItem, exists := r.stock[item.Sku]
		if !exists {
			return fmt.Errorf("product with sku %d not found", item.Sku)
		}
		if stockItem.TotalCount < item.Quantity {
			return fmt.Errorf("insufficient stock for product with sku %d", item.Sku)
		}
		if stockItem.Reserved < item.Quantity {
			return fmt.Errorf("insufficient reserve for product with sku %d", item.Sku)
		}
	}

	for _, item := range items {
		r.stock[item.Sku].TotalCount -= item.Quantity
		r.stock[item.Sku].Reserved -= item.Quantity
	}

	return nil
}
