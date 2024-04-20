package productservice

import (
	"context"

	"route256.ozon.ru/project/cart/internal/pkg/redis"
)

type Client interface {
	GetProduct(ctx context.Context, skuID int64) (*GetProductResponse, error)
	GetProductWithRetries(ctx context.Context, skuID int64) (*GetProductResponse, error)
}

type client struct {
	redisClient redis.Client
}

const (
	ServiceName = "product-service"
)

func NewClient(
	redisClient redis.Client,
) Client {
	return &client{
		redisClient: redisClient,
	}
}
