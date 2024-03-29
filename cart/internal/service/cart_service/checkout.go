package cartservice

import (
	"context"
	"fmt"

	lomsservice "route256.ozon.ru/project/cart/internal/pkg/client/loms_service"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
)

func (s *service) Checkout(ctx context.Context, userID int64) (int64, error) {

	cart, err := s.cartStorage.GetItemsByUserID(ctx, userID)
	if err != nil {
		logger.Error("cartStorage.GetItemsByUserID: failed to get cart items by userID", err)
		return 0, fmt.Errorf("failed to get cart items by userID: %w", err)
	}

	orderID, err := s.lomsService.OrderCreate(ctx, userID, ToOrderItems(cart.Items))
	if err != nil {
		logger.Error("lomsService.OrderCreate: failed to create order", err)
		return 0, fmt.Errorf("failed to create order: %w", err)
	}

	err = s.cartStorage.DeleteItemsByUserID(ctx, userID)
	if err != nil {
		logger.Error("cartStorage.DeleteItemsByUserID: failed to delete cart items by userID", err)
		return 0, fmt.Errorf("failed to delete cart items by userID: %w", err)
	}

	return orderID, nil
}

func ToOrderItems(items cartstorage.CartItems) lomsservice.OrderItems {
	res := make(lomsservice.OrderItems, 0, len(items))
	for sku, item := range items {
		res = append(res, &lomsservice.OrderItem{
			Sku:      sku,
			Quantity: item.Quantity,
		})
	}
	return res
}
