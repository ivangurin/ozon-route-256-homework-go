package cartstorage

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (cs *cartStorage) GetItemsByUserID(
	ctx context.Context,
	userID int64,
) (*Cart, error) {
	cs.RLock()
	defer cs.RUnlock()

	logger.Info(fmt.Sprintf("start get cart for userID: %d", userID))
	defer logger.Info(fmt.Sprintf("finish get cart for userID: %d", userID))

	cart, exists := storage[userID]
	if !exists {
		logger.Info(fmt.Sprintf("cart for userID %d not found", userID))
		return nil, model.ErrNotFound
	}

	return cart, nil
}
