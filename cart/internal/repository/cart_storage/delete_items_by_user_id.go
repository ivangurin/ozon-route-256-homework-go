package cartstorage

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (cs *cartStorage) DeleteItemsByUserID(
	ctx context.Context,
	userID int64,
) error {

	logger.Info(fmt.Sprintf("cartStorage: start clear cart for userID: %d", userID))
	defer logger.Info(fmt.Sprintf("cartStorage: finish clear cart for userID: %d", userID))

	_, exists := storage[userID]
	if exists {
		delete(storage, userID)
	}

	return nil

}
