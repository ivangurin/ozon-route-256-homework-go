package orderservice

import (
	"route256.ozon.ru/project/loms/internal/model"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

func ToOrderStorageItems(items model.OrderItems) orderstorage.OrderItems {
	res := make(orderstorage.OrderItems, 0, len(items))
	for _, item := range items {
		res = append(res, toOrderStorageItem(item))
	}
	return res
}

func toOrderStorageItem(item *model.OrderItem) *orderstorage.OrderItem {
	return &orderstorage.OrderItem{
		Sku:      item.Sku,
		Quantity: item.Quantity,
	}
}

func ToStockItems(items model.OrderItems) stockstorage.ReserveItems {
	res := make(stockstorage.ReserveItems, 0, len(items))
	for _, item := range items {
		res = append(res, toStockItem(item))
	}
	return res
}

func toStockItem(item *model.OrderItem) *stockstorage.ReserveItem {
	return &stockstorage.ReserveItem{
		Sku:      item.Sku,
		Quantity: item.Quantity,
	}
}

func toModelOrder(order *orderstorage.Order) *model.Order {
	return &model.Order{
		ID:     order.ID,
		User:   order.User,
		Status: order.Status,
		Items:  toModelOrderItems(order.Items),
	}
}

func toModelOrderItems(items orderstorage.OrderItems) model.OrderItems {
	res := make(model.OrderItems, 0, len(items))
	for _, item := range items {
		res = append(res, toModelOrderItem(item))
	}
	return res
}

func toModelOrderItem(item *orderstorage.OrderItem) *model.OrderItem {
	return &model.OrderItem{
		Sku:      item.Sku,
		Quantity: item.Quantity,
	}
}
