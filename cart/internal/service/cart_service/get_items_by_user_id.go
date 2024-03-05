package cartservice

import (
	"context"
	"fmt"
	"sort"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
)

func (s *cartService) GetItemsByUserID(ctx context.Context, userID int64) (*Cart, error) {
	logger.Info(fmt.Sprintf("cartService.GetItemsByUserID: start get cart userID: %d", userID))
	defer logger.Info(fmt.Sprintf("cartService.GetItemsByUserID: finish get cart userID: %d", userID))

	cart, err := s.cartStorage.GetItemsByUserID(ctx, userID)
	if err != nil {
		logger.Error("cartService.GetItemsByUserID: faild to get items by user id", err)
		return nil, fmt.Errorf("faild to get items by user id: %w", err)
	}

	resp, err := s.toGetCartResponse(ctx, cart)
	if err != nil {
		logger.Error("cartService.GetItemsByUserID: faild to make response", err)
		return nil, fmt.Errorf("faild to make response: %w", err)
	}

	return resp, nil
}

func (s *cartService) toGetCartResponse(ctx context.Context, cart *cartstorage.Cart) (*Cart, error) {
	resp := &Cart{}
	resp.Items = make([]*CartItem, 0, len(cart.Items))
	for skuID, cartItem := range cart.Items {

		product, err := s.productService.GetProductWithRetries(ctx, skuID)
		if err != nil {
			logger.Error(fmt.Sprintf("cartService.AddItem: faild to get product %d", skuID), err)
			return nil, fmt.Errorf("faild to get product %d: %w", skuID, err)
		}

		resp.Items = append(resp.Items, &CartItem{
			SkuID:    skuID,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: cartItem.Quantity,
		})
		resp.TotalPrice += uint32(cartItem.Quantity) * product.Price
	}

	sort.SliceStable(resp.Items, func(i, j int) bool {
		return resp.Items[i].SkuID < resp.Items[j].SkuID
	})

	return resp, nil
}
