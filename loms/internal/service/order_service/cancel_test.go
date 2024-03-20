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

func TestOrderCancel(t *testing.T) {

	type test struct {
		Name               string
		OrderID            int64
		Order              *orderstorage.Order
		Status             string
		Error              error
		GetByIDError       error
		CancelReserveError error
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
			Order: &orderstorage.Order{
				ID:     2,
				User:   2,
				Status: model.OrederStatusNew,
				Items: []*orderstorage.OrderItem{
					{
						Sku:      2,
						Quantity: 2,
					},
				},
			},
			CancelReserveError: err1,
			Error:              err1,
		},
		{
			Name:    "Ошибка при изменении статуса",
			OrderID: 3,
			Status:  model.OrederStatusCanceled,
			Order: &orderstorage.Order{
				ID:     3,
				User:   3,
				Status: model.OrederStatusNew,
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
			Name:    "Отмена без ошибок",
			OrderID: 4,
			Status:  model.OrederStatusCanceled,
			Order: &orderstorage.Order{
				ID:     4,
				User:   4,
				Status: model.OrederStatusNew,
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
		sp.GetOrderStorege(),
	)

	for _, test := range tests {

		modelOrder := orderservice.ToModelOrder(test.Order)

		sp.GetOrderStoregeMock().GetByIDMock.
			When(test.OrderID).
			Then(test.Order, test.GetByIDError)

		if modelOrder != nil {
			sp.GetStockStoregeMock().CancelReserveMock.
				When(orderservice.ToStockItems(modelOrder.Items)).
				Then(test.CancelReserveError)
		}

		sp.GetOrderStoregeMock().SetStatusMock.
			When(test.OrderID, test.Status).
			Then(test.SetStatusError)

		t.Run(test.Name, func(t *testing.T) {

			err := orderService.Cancel(test.OrderID)
			if test.Error != nil {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не та ошибка")
			} else {
				require.Nil(t, err, "Не должно быть ошибки")
			}

		})

	}
}
