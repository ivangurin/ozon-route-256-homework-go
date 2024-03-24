package lomsservice

type OrderItem struct {
	Sku      int64
	Quantity uint16
}

type OrderItems []*OrderItem
