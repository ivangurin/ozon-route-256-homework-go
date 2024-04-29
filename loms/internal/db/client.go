package db

import (
	"context"
	"fmt"
	"sync/atomic"

	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

type Client interface {
	AddShard(masterUrl, syncUrl string) error
	GetShards() []*shard
	GetShardByUserID(id int64) int64
	GetShardByOrderID(id int64) int64
	GetReaderPoolByShardID(id int64) Pool
	GetReaderPoolByUserID(id int64) Pool
	GetReaderPoolByOrderID(id int64) Pool
	GetWriterPoolByShardID(id int64) Pool
	GetWriterPoolByUserID(id int64) Pool
	GetWriterPoolByOrderID(id int64) Pool
	Close() error
}

type shard struct {
	Master Pool
	Sync   Pool
}

type client struct {
	ctx               context.Context
	shards            []*shard
	readerPoolCounter atomic.Uint64
}

func NewClient(ctx context.Context) Client {
	return &client{
		ctx: ctx,
	}
}

func (c *client) AddShard(masterUrl, syncUrl string) error {

	shard := &shard{}
	c.shards = append(c.shards, shard)

	var err error
	shard.Master, err = NewPool(c.ctx, masterUrl)
	if err != nil {
		return fmt.Errorf("failed to create master pool: %v", err)
	}

	shard.Sync, err = NewPool(c.ctx, syncUrl)
	if err != nil {
		return fmt.Errorf("failed to create sync pool: %v", err)
	}

	return nil
}

func (c *client) GetShardByUserID(id int64) int64 {
	return id % int64(len(c.shards))
}

func (c *client) GetShardByOrderID(id int64) int64 {
	return id % 10
}

func (c *client) GetShards() []*shard {
	return c.shards
}

func (c *client) GetReaderPoolByShardID(id int64) Pool {
	res := c.readerPoolCounter.Add(1)
	if res%2 == 0 {
		return c.GetMasterPoolByShardID(id)
	}
	return c.GetSyncPoolByShardID(id)
}

func (c *client) GetReaderPoolByUserID(id int64) Pool {
	return c.GetReaderPoolByShardID(c.GetShardByUserID(id))
}

func (c *client) GetReaderPoolByOrderID(id int64) Pool {
	return c.GetReaderPoolByShardID(c.GetShardByOrderID(id))
}

func (c *client) GetWriterPoolByShardID(id int64) Pool {
	return c.GetMasterPoolByShardID(id)
}

func (c *client) GetWriterPoolByUserID(id int64) Pool {
	return c.GetWriterPoolByShardID(c.GetShardByUserID(id))
}

func (c *client) GetWriterPoolByOrderID(id int64) Pool {
	return c.GetWriterPoolByShardID(c.GetShardByOrderID(id))
}

func (c *client) GetMasterPoolByShardID(id int64) Pool {
	return c.shards[id].Master
}

func (c *client) GetSyncPoolByShardID(id int64) Pool {
	return c.shards[id].Sync
}

func (c *client) Close() error {
	var err error
	for _, shard := range c.shards {
		if shard.Master != nil {
			err = shard.Master.Close()
			if err != nil {
				logger.Errorf(c.ctx, "failed to close master pool: %v", err)
			}
		}
		if shard.Sync != nil {
			err = shard.Sync.Close()
			if err != nil {
				logger.Errorf(c.ctx, "failed to close sync pool: %v", err)
			}
		}
	}
	return nil
}
