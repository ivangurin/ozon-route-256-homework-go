package cartservice

import (
	"context"
	"fmt"
	"sort"

	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func (s *service) toGetCartResponse(ctx context.Context, cart *cartstorage.Cart) (*Cart, error) {
	var resp *Cart = &Cart{}
	resp.Items = make([]*CartItem, 0, len(cart.Items))
	for skuID, cartItem := range cart.Items {

		product, err := s.productService.GetProductWithRetries(ctx, skuID)
		if err != nil {
			logger.Errorf(ctx, "cartService.AddItem: failed to get product %d: %v", skuID, err)
			return nil, fmt.Errorf("failed to get product %d: %w", skuID, err)
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
