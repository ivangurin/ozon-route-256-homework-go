package notifierservice

import (
	"context"
)

type Service interface {
	OrderStatusChanges(ctx context.Context, orderID int64, status string) error
}

type service struct {
}

func NewService() Service {
	return &service{}
}
