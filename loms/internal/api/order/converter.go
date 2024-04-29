package orderapi

import (
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/pkg/api/order/v1"
)

func toItems(items []*order.OrderCreateRequest_Item) model.OrderItems {
	res := make(model.OrderItems, 0, len(items))
	for _, item := range items {
		res = append(res, toItem(item))
	}
	return res
}

func toItem(item *order.OrderCreateRequest_Item) *model.OrderItem {
	return &model.OrderItem{
		Sku:      item.GetSku(),
		Quantity: uint16(item.GetCount()),
	}
}

func toRespOrders(orders []*model.Order) []*order.Order {
	res := make([]*order.Order, 0, len(orders))
	for _, order := range orders {
		res = append(res, toRespOrder(order))
	}
	return res
}

func toRespOrder(ord *model.Order) *order.Order {
	return &order.Order{
		User:   ord.User,
		Status: ord.Status,
		Items:  toRespItems(ord.Items),
	}
}

func toRespItems(items model.OrderItems) []*order.OrderItem {
	res := make([]*order.OrderItem, 0, len(items))
	for _, item := range items {
		res = append(res, toRespItem(item))
	}
	return res
}

func toRespItem(item *model.OrderItem) *order.OrderItem {
	return &order.OrderItem{
		Sku:   item.Sku,
		Count: uint64(item.Quantity),
	}
}
