package orderservice

import "time"

const (
	OrderEventStatusChanged = "order_status_changed"
)

const (
	OrderEntityOrder = "order"
)

type GenericMessage struct {
	Event string `json:"event"`
}

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
