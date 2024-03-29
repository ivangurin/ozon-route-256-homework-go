package stockstorage

import "route256.ozon.ru/project/loms/internal/model"

func (r *repository) GetBySku(sku int64) (uint16, error) {
	r.RLock()
	defer r.RUnlock()

	stockItem, exists := r.stock[sku]
	if !exists {
		return 0, model.ErrNotFound
	}

	return stockItem.TotalCount - stockItem.Reserved, nil
}
