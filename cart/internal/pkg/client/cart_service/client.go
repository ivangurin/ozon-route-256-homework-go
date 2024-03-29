package cartservice

import "context"

type Client interface {
	AddItem(ctx context.Context, UserID int64, SkuID int64, Quantity uint16) error
	DeleteItem(ctx context.Context, UserID int64, SkuID int64) error
	DeleteItemsByUserID(ctx context.Context, UserID int64) error
	GetItemsByUserID(ctx context.Context, UserID int64) (*GetItmesByUserIDResponse, error)
}

type client struct {
	host string
}

func NewClient(host string) Client {
	return &client{
		host: host,
	}
}
