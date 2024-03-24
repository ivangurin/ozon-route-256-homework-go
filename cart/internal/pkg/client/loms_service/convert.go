package lomsservice

import (
	order_api "route256.ozon.ru/project/cart/internal/pb/api/order/v1"
)

func ToOrderCreateRequest(user int64, items OrderItems) *order_api.OrderCreateRequest {
	return &order_api.OrderCreateRequest{
		User:  user,
		Items: toOrderCreateRequestItems(items),
	}
}

func toOrderCreateRequestItems(items OrderItems) []*order_api.OrderCreateRequest_Item {
	res := make([]*order_api.OrderCreateRequest_Item, 0, len(items))
	for _, item := range items {
		res = append(res, toOrderCreateRequestItem(item))
	}
	return res
}

func toOrderCreateRequestItem(item *OrderItem) *order_api.OrderCreateRequest_Item {
	return &order_api.OrderCreateRequest_Item{
		Sku:   item.Sku,
		Count: uint64(item.Quantity),
	}
}
