package cartservice

type AddItemRequestBody struct {
	Count uint16 `json:"count"`
}

type GetItmesByUserIDResponse struct {
	Items      []*GetItmesByUserIDResponseItem `json:"items"`
	TotalPrice uint32                          `json:"total_price"`
}

type GetItmesByUserIDResponseItem struct {
	SkuID    int64  `json:"sku_id"`
	Name     string `json:"name"`
	Quantity uint16 `json:"count"`
	Price    uint32 `json:"price"`
}
