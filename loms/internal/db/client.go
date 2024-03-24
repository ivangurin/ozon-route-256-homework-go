package db

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/loms/internal/config"
)

type Client interface {
	GetMasterPool() Pool
	GetSyncPool() Pool
	Close() error
}

type client struct {
	ctx        context.Context
	masterPool Pool
	syncPool   Pool
}

func NewClient(ctx context.Context) (Client, error) {
	masterPoll, err := NewPool(ctx, config.MasterDBUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create master pool: %v", err)
	}

	syncPoll, err := NewPool(ctx, config.SyncDBUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create sync pool: %v", err)
	}

	return &client{
		ctx:        ctx,
		masterPool: masterPoll,
		syncPool:   syncPoll,
	}, nil
}

func (c *client) GetMasterPool() Pool {
	return c.masterPool
}

func (c *client) GetSyncPool() Pool {
	return c.syncPool
}

func (c *client) Close() error {
	c.masterPool.Close()
	c.syncPool.Close()
	return nil
}
