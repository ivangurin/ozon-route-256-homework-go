package cartstorage

import (
	"context"
	"time"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
	"route256.ozon.ru/project/cart/internal/pkg/metrics"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func (s *storage) DeleteItemsByUserID(
	ctx context.Context,
	userID int64,
) error {
	_, span := tracer.StartSpanFromContext(ctx, "cartStorage.DeleteItemsByUserID")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"DeleteItemsByUserID",
		"delete",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	s.Lock()
	defer s.Unlock()

	logger.Infof(ctx, "cartStorage: start clear cart for userID: %d", userID)
	defer logger.Infof(ctx, "cartStorage: finish clear cart for userID: %d", userID)

	_, exists := cartStorage[userID]
	if exists {
		delete(cartStorage, userID)
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"DeleteItemsByUserID",
		"delete",
		"ok",
	)
	return nil
}
