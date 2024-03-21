package cartservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *service) GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error) {
	cart, err := s.cartStorage.GetItemsByUserID(ctx, userID)
	if err != nil {
		logger.Error("cartService.GetItemsByUserID: failed to get items by user id", err)
		return nil, fmt.Errorf("failed to get items by user id: %w", err)
	}

	resp, err := s.toGetCartResponse(ctx, cart)
	if err != nil {
		logger.Error("cartService.GetItemsByUserID: failed to make response", err)
		return nil, fmt.Errorf("failed to make response: %w", err)
	}

	return resp, nil
}
