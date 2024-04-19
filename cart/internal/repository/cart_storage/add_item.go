package cartstorage

import (
	"context"
	"time"

	"route256.ozon.ru/project/cart/internal/pkg/metrics"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func (s *storage) AddItem(
	ctx context.Context,
	userID int64,
	skuID int64,
	quantity uint16,
) error {
	_, span := tracer.StartSpanFromContext(ctx, "cartStorage.AddItem")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"AddItem",
		"insert",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	s.Lock()
	defer s.Unlock()

	cart, exists := cartStorage[userID]
	if !exists {
		cart = &Cart{
			Items: CartItems{},
		}
		cartStorage[userID] = cart
	}

	cartItem, exists := cart.Items[skuID]
	if !exists {
		cartItem = &CartItem{}
		cart.Items[skuID] = cartItem
	}

	cartItem.Quantity += quantity

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"AddItem",
		"insert",
		"ok",
	)
	return nil
}
