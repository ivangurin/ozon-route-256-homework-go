package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"route256.ozon.ru/project/cart/internal/config"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

type Client interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, value interface{}) (bool, error)
	Close() error
}

type client struct {
	client *redis.Client
}

func NewClient() (Client, error) {
	opts, err := redis.ParseURL(config.RedisUrl)
	if err != nil {
		return nil, err
	}

	return &client{
		client: redis.NewClient(opts),
	}, nil
}

func (c *client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	ctx, span := tracer.StartSpanFromContext(ctx, fmt.Sprintf("redisClient.Set/%s", key))
	defer span.End()

	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, json, ttl).Err()
}

func (c *client) Get(ctx context.Context, key string, value interface{}) (bool, error) {
	ctx, span := tracer.StartSpanFromContext(ctx, fmt.Sprintf("redisClient.Get/%s", key))
	defer span.End()

	str, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		} else {
			return false, err
		}
	}

	err = json.Unmarshal([]byte(str), value)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *client) Close() error {
	return c.client.Close()
}
