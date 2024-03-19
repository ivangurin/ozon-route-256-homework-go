package model

const (
	OrederStatusNew            = "new"
	OrederStatusAwatingPayment = "awaiting_payment"
	OrederStatusPayed          = "payed"
	OrederStatusCanceled       = "canceled"
	OrederStatusFailed         = "failed"
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
