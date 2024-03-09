package cartstorage

import "context"

var storage map[int64]*Cart = map[int64]*Cart{}

type IStorage interface {
	AddItem(ctx context.Context, userID int64, skuID int64, quantity uint16) error
	DeleteItem(ctx context.Context, userID int64, skuID int64) error
	GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error)
	DeleteItemsByUserID(ctx context.Context, userID int64) error
}

type cartStorage struct{}

func NewCartStorage() IStorage {
	return &cartStorage{}
}
