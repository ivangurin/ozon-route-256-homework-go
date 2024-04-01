package orderservice_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/suite"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	orderservice "route256.ozon.ru/project/loms/internal/service/order_service"
)

func TestOrderPay(t *testing.T) {

	type test struct {
		Name               string
		OrderID            int64
		Order              *orderstorage.Order
		Status             string
		Error              error
		GetByIDError       error
		RemoveReserveError error
		SetStatusError     error
	}

	tests := []*test{
		{
			Name:         "Заказ не найден",
			OrderID:      1,
			GetByIDError: model.ErrNotFound,
			Error:        model.ErrNotFound,
		},
		{
			Name:    "Ошибка при снятии резерва",
			OrderID: 2,
			Status:  model.OrderStatusPayed,
			Order: &orderstorage.Order{
				ID:     2,
				User:   2,
				Status: model.OrderStatusNew,
				Items: []*orderstorage.OrderItem{
					{
						Sku:      2,
						Quantity: 2,
					},
				},
			},
			RemoveReserveError: err1,
			Error:              err1,
		},
		{
			Name:    "Ошибка при изменении статуса",
			OrderID: 3,
			Status:  model.OrderStatusPayed,
			Order: &orderstorage.Order{
				ID:     3,
				User:   3,
				Status: model.OrderStatusNew,
				Items: []*orderstorage.OrderItem{
					{
						Sku:      3,
						Quantity: 3,
					},
				},
			},
			SetStatusError: err2,
			Error:          err2,
		},
		{
			Name:    "Оплата без ошибок",
			OrderID: 4,
			Status:  model.OrderStatusPayed,
			Order: &orderstorage.Order{
				ID:     4,
				User:   4,
				Status: model.OrderStatusNew,
				Items: []*orderstorage.OrderItem{
					{
						Sku:      4,
						Quantity: 4,
					},
				},
			},
		},
	}

	ctx := context.Background()

	sp := suite.NewSuiteProvider(t, ctx)

	orderService := orderservice.NewService(
		ctx,
		sp.GetStockStorage(),
		sp.GetOrderStorage(),
	)

	for _, test := range tests {

		modelOrder := orderservice.ToModelOrder(test.Order)

		sp.GetOrderStorageMock().GetByIDMock.
			When(ctx, test.OrderID).
			Then(test.Order, test.GetByIDError)

		if modelOrder != nil {
			sp.GetStockStorageMock().RemoveReserveMock.
				When(ctx, orderservice.ToStockItems(modelOrder.Items)).
				Then(test.RemoveReserveError)
		}

		sp.GetOrderStorageMock().SetStatusMock.
			When(ctx, test.OrderID, test.Status).
			Then(test.SetStatusError)

		t.Run(test.Name, func(t *testing.T) {

			err := orderService.Pay(ctx, test.OrderID)
			if test.Error != nil {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не та ошибка")
			} else {
				require.Nil(t, err, "Не должно быть ошибки")
			}

		})

	}
}
