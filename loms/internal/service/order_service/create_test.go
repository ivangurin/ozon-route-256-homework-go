package orderservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/suite"
	orderservice "route256.ozon.ru/project/loms/internal/service/order_service"
)

var err1 = errors.New("some error 1")
var err2 = errors.New("some error 2")
var err3 = errors.New("some error 3")
var err4 = errors.New("some error 4")

func TestOrderCreate(t *testing.T) {

	type test struct {
		Name              string
		User              int64
		Items             model.OrderItems
		OrderID           int64
		Status            string
		Error             error
		CreateOrderError  error
		ReserveStockError error
		SetStatusError    error
	}

	tests := []*test{
		{
			Name:             "Ощибка при создании заказа",
			Items:            model.OrderItems{&model.OrderItem{Sku: 1, Quantity: 1}},
			Error:            err1,
			CreateOrderError: err1,
		},
		{
			Name:              "Ошибка при изменении статуса при не успешном резервировании",
			Items:             model.OrderItems{&model.OrderItem{Sku: 1, Quantity: 1}},
			Status:            model.OrederStatusFailed,
			Error:             err3,
			ReserveStockError: err2,
			SetStatusError:    err3,
		},
		{
			Name:           "Ошибка при изменении статуса при успешном резервировании",
			Items:          model.OrderItems{&model.OrderItem{Sku: 1, Quantity: 1}},
			Status:         model.OrederStatusAwatingPayment,
			Error:          err4,
			SetStatusError: err4,
		},
		{
			Name:    "Успешное создание заказа",
			Items:   model.OrderItems{&model.OrderItem{Sku: 1, Quantity: 1}},
			Status:  model.OrederStatusAwatingPayment,
			OrderID: 1,
		},
	}

	ctx := context.Background()

	for _, test := range tests {

		sp := suite.NewSuiteProvider(t, ctx)

		orderService := orderservice.NewService(
			ctx,
			sp.GetStockStorage(),
			sp.GetOrderStorege(),
		)

		sp.GetOrderStoregeMock().CreateMock.
			When(test.User, orderservice.ToOrderStorageItems(test.Items)).
			Then(test.OrderID, test.CreateOrderError)

		sp.GetStockStoregeMock().ReserveMock.
			When(orderservice.ToStockItems(test.Items)).
			Then(test.ReserveStockError)

		sp.GetOrderStoregeMock().SetStatusMock.
			When(test.OrderID, test.Status).
			Then(test.SetStatusError)

		t.Run(test.Name, func(t *testing.T) {

			orderID, err := orderService.Create(test.User, test.Items)
			if test.Error != nil {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не та ошибка")
			} else {
				require.Equal(t, test.OrderID, orderID, "Не совпало количество")
			}

		})
	}

}
