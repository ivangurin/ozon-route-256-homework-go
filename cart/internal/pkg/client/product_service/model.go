package productservice

type GetProductRequest struct {
	Token string `json:"token"`
	Sku   int64  `json:"sku"`
}

type GetProductResponse struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}
