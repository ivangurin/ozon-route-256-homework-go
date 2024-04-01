package cartstorage

import (
	"context"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *storage) DeleteItemsByUserID(
	ctx context.Context,
	userID int64,
) error {
	s.Lock()
	defer s.Unlock()

	logger.Infof("cartStorage: start clear cart for userID: %d", userID)
	defer logger.Infof("cartStorage: finish clear cart for userID: %d", userID)

	_, exists := cartStorage[userID]
	if exists {
		delete(cartStorage, userID)
	}

	return nil
}
