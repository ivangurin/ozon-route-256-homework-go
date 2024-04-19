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
