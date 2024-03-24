package stockstorage

type StockItem struct {
	Sku        int64  `json:"sku"`
	TotalCount uint16 `json:"total_count"`
	Reserved   uint16 `json:"reserved"`
}

type Stock map[int64]*StockItem

type ReserveItem struct {
	Sku      int64
	Quantity uint16
}

type ReserveItems []*ReserveItem
