package orderservice_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/model"
	"route256.ozon.ru/project/loms/internal/pkg/suite"
	orderservice "route256.ozon.ru/project/loms/internal/service/order_service"
)

var (
	err1 = errors.New("some error 1")
	err2 = errors.New("some error 2")
	err3 = errors.New("some error 3")
	err4 = errors.New("some error 4")
)

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
			Name:             "Ошибка при создании заказа",
			User:             1,
			Items:            model.OrderItems{&model.OrderItem{Sku: 1, Quantity: 1}},
			Error:            err1,
			CreateOrderError: err1,
		},
		{
			Name:              "Ошибка при изменении статуса при не успешном резервировании",
			User:              2,
			Items:             model.OrderItems{&model.OrderItem{Sku: 1, Quantity: 1}},
			Status:            model.OrderStatusFailed,
			Error:             err3,
			ReserveStockError: err2,
			SetStatusError:    err3,
		},
		{
			Name:           "Ошибка при изменении статуса при успешном резервировании",
			User:           3,
			Items:          model.OrderItems{&model.OrderItem{Sku: 1, Quantity: 1}},
			Status:         model.OrderStatusAwaitingPayment,
			Error:          err4,
			SetStatusError: err4,
		},
		{
			Name:    "Успешное создание заказа",
			User:    4,
			Items:   model.OrderItems{&model.OrderItem{Sku: 1, Quantity: 1}},
			Status:  model.OrderStatusAwaitingPayment,
			OrderID: 1,
		},
	}

	t.Parallel()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()

			orderService := orderservice.NewService(
				sp.GetStockStorage(),
				sp.GetOrderStorage(),
				sp.GetKafkaProducer(),
			)

			sp.GetOrderStorageMock().EXPECT().
				Create(mock.Anything, test.User, mock.Anything).
				Return(test.OrderID, test.CreateOrderError)

			sp.GetStockStorageMock().EXPECT().
				Reserve(mock.Anything, orderservice.ToStockItems(test.Items)).
				Return(test.ReserveStockError)

			sp.GetOrderStorageMock().EXPECT().
				SetStatus(mock.Anything, test.OrderID, test.Status).
				Return(test.SetStatusError)

			sp.GetKafkaProducer().EXPECT().
				SendMessageWithKey(mock.Anything, config.KafkaOrderEventsTopic, fmt.Sprintf("%d", test.OrderID), mock.Anything).
				Return(nil)

			orderID, err := orderService.Create(context.Background(), test.User, test.Items)
			if test.Error != nil {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не та ошибка")
			} else {
				require.Nil(t, err, "Не должно быть ошибки")
				require.Equal(t, test.OrderID, orderID, "Не совпало количество")
			}
		})
	}

}
