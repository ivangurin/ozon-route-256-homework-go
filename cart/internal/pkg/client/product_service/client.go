package productservice

import (
	"context"
)

type Client interface {
	GetProduct(ctx context.Context, skuID int64) (*GetProductResponse, error)
	GetProductWithRetries(ctx context.Context, skuID int64) (*GetProductResponse, error)
}

type client struct{}

const (
	ServiceName = "product-service"
)

var productStorage map[int64]*GetProductResponse = map[int64]*GetProductResponse{}

func NewClient() Client {
	return &client{}
}
