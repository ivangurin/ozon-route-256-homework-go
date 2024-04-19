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
	EventOrderStatusChanged = "order_status_changed"
)

const (
	EntityTypeOrder = "order"
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
	ID         string                       `json:"id"`
	Time       time.Time                    `json:"time"`
	Event      string                       `json:"event"`
	EntityType string                       `json:"entity_type"`
	EntityID   string                       `json:"entity_id"`
	Data       OrderChangeStatusMessageData `json:"data"`
}

type OrderChangeStatusMessageData struct {
	Order OrderChangeStatusMessageOrder `json:"order"`
}

type OrderChangeStatusMessageOrder struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}
