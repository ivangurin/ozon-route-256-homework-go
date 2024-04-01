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
	ctx               context.Context
	masterPool        Pool
	syncPool          Pool
	readerPoolCounter atomic.Uint64
}

func NewClient(ctx context.Context, masterDBUrl, syncDBUrl string) (Client, error) {
	masterPool, err := NewPool(ctx, masterDBUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create master pool: %v", err)
	}

	syncPool, err := NewPool(ctx, syncDBUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create sync pool: %v", err)
	}

	return &client{
		ctx:        ctx,
		masterPool: masterPool,
		syncPool:   syncPool,
	}, nil
}

func (c *client) GetReaderPool() Pool {
	res := c.readerPoolCounter.Add(1)
	if res%2 == 0 {
		return c.masterPool
	}
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
