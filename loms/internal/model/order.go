package model

import "time"

const (
	OrderStatusNew             = "new"
	OrderStatusAwaitingPayment = "awaiting_payment"
	OrderStatusPayed           = "payed"
	OrderStatusCancelled       = "cancelled"
	OrderStatusFailed          = "failed"
)

const (
	OrderEventStatusChanged = "order_status_changed"
)

const (
	OrderEntityOrder = "order"
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

type OrderChangeStatusMessage struct {
	Event  string                       `json:"event"`
	Entity string                       `json:"entity"`
	ID     string                       `json:"id"`
	UUID   string                       `json:"uuid"`
	Time   time.Time                    `json:"time"`
	Data   OrderChangeStatusMessageData `json:"data"`
}

type OrderChangeStatusMessageData struct {
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
}
