package cartapi

type AddItemRequest struct {
	UserID   int64  `validate:"required|int|min:1"`
	SkuID    int64  `validate:"required|int|min:1"`
	Quantity uint16 `validate:"required|int|min:1"`
}

type AddItemRequestBody struct {
	Count uint16 `json:"count"`
}

type DeleteItemRequest struct {
	UserID int64 `validate:"required|int|min:1"`
	SkuID  int64 `validate:"required|int|min:1"`
}

type GetItemsByUserIDRequest struct {
	UserID int64 `validate:"required|int|min:1"`
}

type DeleteItemsByUserIDRequest struct {
	UserID int64 `validate:"required|int|min:1"`
}

type CheckoutRequest struct {
	UserID int64 `json:"user" validate:"required|int|min:1"`
}

type CheckoutResponse struct {
	OrderID int64 `json:"order_id"`
}
