package db

import (
	"context"
	"fmt"
	"sync/atomic"
)

type Client interface {
	GetReaderPool() Pool
	GetWriterPool() Pool
	GetMasterPool() Pool
	GetSyncPool() Pool
	Close() error
}

type client struct {
	ctx        context.Context
	masterPool Pool
	syncPool   Pool
	readerPool atomic.Int32
}

func NewClient(ctx context.Context, masterDBUrl, syncDBUrl string) (Client, error) {
	masterPoll, err := NewPool(ctx, masterDBUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create master pool: %v", err)
	}

	syncPoll, err := NewPool(ctx, syncDBUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create sync pool: %v", err)
	}

	return &client{
		ctx:        ctx,
		masterPool: masterPoll,
		syncPool:   syncPoll,
	}, nil
}

func (c *client) GetReaderPool() Pool {
	if c.readerPool.Load() == 0 {
		c.readerPool.Add(1)
		return c.masterPool
	}
	c.readerPool.Add(-1)
	return c.syncPool
}

func (c *client) GetWriterPool() Pool {
	return c.GetMasterPool()
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
