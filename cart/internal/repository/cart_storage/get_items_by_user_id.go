package cartstorage

import (
	"context"
	"time"

	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
	"route256.ozon.ru/project/cart/internal/pkg/metrics"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func (s *storage) GetItemsByUserID(
	ctx context.Context,
	userID int64,
) (*Cart, error) {
	_, span := tracer.StartSpanFromContext(ctx, "cartStorage.GetItemsByUserID")
	defer span.End()

	metrics.UpdateDatabaseRequestsTotal(
		RepositoryName,
		"GetItemsByUserID",
		"select",
	)
	defer metrics.UpdateDatabaseResponseTime(time.Now().UTC())

	s.RLock()
	defer s.RUnlock()

	logger.Infof(ctx, "start get cart for userID: %d", userID)
	defer logger.Infof(ctx, "finish get cart for userID: %d", userID)

	cart, exists := cartStorage[userID]
	if !exists {
		logger.Infof(ctx, "cart for userID %d not found", userID)
		metrics.UpdateDatabaseResponseCode(
			RepositoryName,
			"GetItemsByUserID",
			"select",
			"not_found",
		)
		return nil, model.ErrNotFound
	}

	metrics.UpdateDatabaseResponseCode(
		RepositoryName,
		"GetItemsByUserID",
		"select",
		"ok",
	)
	return cart, nil
}
