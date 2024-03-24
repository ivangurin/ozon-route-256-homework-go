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

func toRespOrder(ord *model.Order) *order.OrderInfoResponse {
	return &order.OrderInfoResponse{
		User:   ord.User,
		Status: ord.Status,
		Items:  toRespItems(ord.Items),
	}
}

func toRespItems(items model.OrderItems) []*order.OrderInfoResponse_Item {
	res := make([]*order.OrderInfoResponse_Item, 0, len(items))
	for _, item := range items {
		res = append(res, toRespItem(item))
	}
	return res
}

func toRespItem(item *model.OrderItem) *order.OrderInfoResponse_Item {
	return &order.OrderInfoResponse_Item{
		Sku:   item.Sku,
		Count: uint64(item.Quantity),
	}
}
