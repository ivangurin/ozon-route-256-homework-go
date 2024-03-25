package cartstorage

import (
	"context"

	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *storage) GetItemsByUserID(
	ctx context.Context,
	userID int64,
) (*Cart, error) {
	s.RLock()
	defer s.RUnlock()

	logger.Infof(ctx, "start get cart for userID: %d", userID)
	defer logger.Infof(ctx, "finish get cart for userID: %d", userID)

	cart, exists := cartStorage[userID]
	if !exists {
		logger.Infof(ctx, "cart for userID %d not found", userID)
		return nil, model.ErrNotFound
	}

	return cart, nil
}
