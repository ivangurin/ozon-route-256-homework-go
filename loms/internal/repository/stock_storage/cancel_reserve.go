package stockstorage

import "fmt"

func (r *repository) CancelReserve(items ReserveItems) error {
	r.Lock()
	defer r.Unlock()

	for _, item := range items {
		stockItem, exists := r.stock[item.Sku]
		if !exists {
			return fmt.Errorf("product with sku %d not found", item.Sku)
		}
		if stockItem.Reserved < item.Quantity {
			return fmt.Errorf("insufficient reseve for product with sku %d", item.Sku)
		}
	}

	for _, item := range items {
		r.stock[item.Sku].Reserved -= item.Quantity
	}

	return nil
}
