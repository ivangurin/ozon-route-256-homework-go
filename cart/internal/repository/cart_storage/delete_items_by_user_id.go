package cartstorage

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *storage) DeleteItemsByUserID(
	ctx context.Context,
	userID int64,
) error {
	s.Lock()
	defer s.Unlock()

	logger.Info(fmt.Sprintf("cartStorage: start clear cart for userID: %d", userID))
	defer logger.Info(fmt.Sprintf("cartStorage: finish clear cart for userID: %d", userID))

	_, exists := cartStorage[userID]
	if exists {
		delete(cartStorage, userID)
	}

	return nil
}
