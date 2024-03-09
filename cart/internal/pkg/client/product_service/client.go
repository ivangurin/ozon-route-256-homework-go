package productservice

import (
	"context"
)

type IClient interface {
	GetProduct(ctx context.Context, skuID int64) (*GetProductResponse, error)
	GetProductWithRetries(ctx context.Context, skuID int64) (*GetProductResponse, error)
}

type client struct{}

var productStorage map[int64]*GetProductResponse = map[int64]*GetProductResponse{}

func NewClient() IClient {
	return &client{}
}
