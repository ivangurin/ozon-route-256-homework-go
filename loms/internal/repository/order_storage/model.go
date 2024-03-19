package orderstorage

type OrderItem struct {
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

type Orders map[int64]*Order

func (r *repository) getNextID() int64 {
	var maxID int64
	for orderID := range r.orders {
		if orderID > maxID {
			maxID = orderID
		}
	}
	maxID++
	return maxID
}
