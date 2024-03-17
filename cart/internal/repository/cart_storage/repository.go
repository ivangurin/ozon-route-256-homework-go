package cartstorage

import (
	"context"
	"sync"
)

var storage map[int64]*Cart = map[int64]*Cart{}

type IStorage interface {
	AddItem(ctx context.Context, userID int64, skuID int64, quantity uint16) error
	DeleteItem(ctx context.Context, userID int64, skuID int64) error
	GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error)
	DeleteItemsByUserID(ctx context.Context, userID int64) error
	Reset()
}

type cartStorage struct {
	sync.RWMutex
}

func NewCartStorage() IStorage {
	return &cartStorage{}
}
