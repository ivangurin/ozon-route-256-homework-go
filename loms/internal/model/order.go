package model

const (
	OrderStatusNew             = "new"
	OrderStatusAwaitingPayment = "awaiting_payment"
	OrderStatusPayed           = "payed"
	OrderStatusCanceled        = "canceled"
	OrderStatusFailed          = "failed"
)

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
