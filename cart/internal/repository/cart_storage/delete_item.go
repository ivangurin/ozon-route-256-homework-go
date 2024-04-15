package cartstorage

import (
	"context"
	"time"

	"route256.ozon.ru/project/cart/internal/pkg/metrics"
)

func (s *storage) DeleteItem(
	ctx context.Context,
	userID int64,
	skuID int64,
) error {
	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"DeleteItem",
		"delete",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	s.Lock()
	defer s.Unlock()

	cart, exists := cartStorage[userID]
	if !exists {
		return nil
	}

	_, exists = cart.Items[skuID]
	if !exists {
		return nil
	}

	delete(cart.Items, skuID)

	if len(cart.Items) == 0 {
		delete(cartStorage, userID)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"DeleteItem",
		"delete",
		"ok",
	)
	return nil
}
