package cartstorage

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *storage) GetItemsByUserID(
	ctx context.Context,
	userID int64,
) (*Cart, error) {
	s.RLock()
	defer s.RUnlock()

	logger.Info(fmt.Sprintf("start get cart for userID: %d", userID))
	defer logger.Info(fmt.Sprintf("finish get cart for userID: %d", userID))

	cart, exists := cartStorage[userID]
	if !exists {
		logger.Info(fmt.Sprintf("cart for userID %d not found", userID))
		return nil, model.ErrNotFound
	}

	return cart, nil
}
