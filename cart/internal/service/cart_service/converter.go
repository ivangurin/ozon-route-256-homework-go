package cartservice

import (
	"context"
	"fmt"
	"sort"

	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"

	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
)

func (s *service) toGetCartResponse(ctx context.Context, cart *cartstorage.Cart, products map[int64]*productservice.GetProductResponse) (*Cart, error) {
	var resp *Cart = &Cart{}
	resp.Items = make([]*CartItem, 0, len(cart.Items))
	for sku, cartItem := range cart.Items {

		product, exists := products[sku]
		if !exists {
			return nil, fmt.Errorf("failed to get product %d", sku)
		}

		resp.Items = append(resp.Items, &CartItem{
			SkuID:    sku,
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
