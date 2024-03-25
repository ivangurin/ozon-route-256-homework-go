package orderstorage

type OrderItem struct {
	ID       int64
	Sku      int64
	Quantity uint16
}

type OrderItems []*OrderItem

type Order struct {
	ID     int64
	User   int64
	Status string
	Items  OrderItems
}
