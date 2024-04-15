package cartstorage

import (
	"context"
	"sync"
)

const (
	RepositoryName = "CartStorage"
)

var cartStorage map[int64]*Cart = map[int64]*Cart{}

type Storage interface {
	AddItem(ctx context.Context, userID int64, skuID int64, quantity uint16) error
	DeleteItem(ctx context.Context, userID int64, skuID int64) error
	GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error)
	DeleteItemsByUserID(ctx context.Context, userID int64) error
	Reset()
}

type storage struct {
	sync.RWMutex
}

func NewCartStorage() Storage {
	return &storage{}
}
