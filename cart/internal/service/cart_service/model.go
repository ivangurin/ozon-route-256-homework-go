package cartservice

type Cart struct {
	Items      []*CartItem `json:"items"`
	TotalPrice uint32      `json:"total_price"`
}

type CartItem struct {
	SkuID    int64  `json:"sku_id"`
	Name     string `json:"name"`
	Quantity uint16 `json:"count"`
	Price    uint32 `json:"price"`
}
