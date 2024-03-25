package stockstorage

type ReserveItem struct {
	Sku      int64
	Quantity uint16
}

type ReserveItems []*ReserveItem
