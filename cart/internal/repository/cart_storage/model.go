package cartstorage

type CartItem struct {
	Quantity uint16
}

type CartItems map[int64]*CartItem

type Cart struct {
	Items CartItems
}
