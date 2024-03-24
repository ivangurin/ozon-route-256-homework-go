package orderservice_test

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/suite"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	orderservice "route256.ozon.ru/project/loms/internal/service/order_service"
)

func TestOrderInfo(t *testing.T) {

	type test struct {
		Name         string
		OrderID      int64
		Order        *orderstorage.Order
		Error        error
		GetByIDError error
	}

	tests := []*test{
		{
			Name:         "Заказ не найден",
			OrderID:      1,
			GetByIDError: model.ErrNotFound,
			Error:        model.ErrNotFound,
		},
		{
			Name:    "Заказ существует",
			OrderID: 2,
			Order: &orderstorage.Order{
				ID:     2,
				User:   1,
				Status: model.OrderStatusNew,
				Items: []*orderstorage.OrderItem{
					{
						Sku:      1,
						Quantity: 1,
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

		sp.GetOrderStorageMock().GetByIDMock.
			When(ctx, test.OrderID).
			Then(test.Order, test.GetByIDError)

		t.Run(test.Name, func(t *testing.T) {

			order, err := orderService.Info(ctx, test.OrderID)
			if test.Error != nil {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не та ошибка")
			} else {
				require.Nil(t, err, "Не должно быть ошибки")
				diff := deep.Equal(orderservice.ToModelOrder(test.Order), order)
				if diff != nil {
					t.Errorf("Заказы должны совпасть: %+v", diff)
				}
			}

		})
	}

}
