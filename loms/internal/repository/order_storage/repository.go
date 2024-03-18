package orderstorage

import "context"

type Repository interface {
}

type repository struct {
}

func NewRepository(ctx context.Context) Repository {
	return &repository{}
}
