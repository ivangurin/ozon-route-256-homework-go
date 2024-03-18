package orderservice

import "context"

type Service interface {
}

type service struct {
}

func NewService(ctx context.Context) Service {
	return &service{}
}
