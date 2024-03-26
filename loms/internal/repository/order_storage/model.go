package orderstorage

type Order struct {
	ID     int64
	User   int64
	Status string
	Items  OrderItems
}

type OrderItem struct {
	ID       int64
	Sku      int64
	Quantity uint16
}

type OrderItems []*OrderItem

func (o *Order) GetItemsMap() map[int64]*OrderItem {
	res := make(map[int64]*OrderItem, len(o.Items))
	for _, item := range o.Items {
		res[item.Sku] = item
	}
	return res
}
