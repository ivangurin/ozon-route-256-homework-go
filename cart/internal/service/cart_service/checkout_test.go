package cartservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/pb/api/order/v1"
	lomsservice "route256.ozon.ru/project/cart/internal/pkg/client/loms_service"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestCheckout(t *testing.T) {

	type test struct {
		Name                     string
		UserID                   int64
		OrderID                  int64
		Cart                     *cartstorage.Cart
		GetItemsByUserIDError    error
		CreateOrderError         error
		DeleteItemsByUserIDerror error
		Error                    error
	}

	var (
		err1 = fmt.Errorf("some error 1")
		err2 = fmt.Errorf("some error 2")
		err3 = fmt.Errorf("some error 3")
	)

	ctx := context.Background()

	tests := []*test{
		{
			Name:                  "Корзина не создана",
			UserID:                1,
			GetItemsByUserIDError: err1,
			Error:                 err1,
		},
		{
			Name:   "Ошибка при создании заказа",
			UserID: 2,
			Cart: &cartstorage.Cart{
				Items: map[int64]*cartstorage.CartItem{
					1: &cartstorage.CartItem{
						Quantity: 1,
					},
				},
			},
			CreateOrderError: err2,
			Error:            err2,
		},
		{
			Name:   "Ошибка при удалении корзины",
			UserID: 3,
			Cart: &cartstorage.Cart{
				Items: map[int64]*cartstorage.CartItem{
					1: &cartstorage.CartItem{
						Quantity: 1,
					},
				},
			},
			DeleteItemsByUserIDerror: err3,
			Error:                    err3,
		},
		{
			Name:   "Успешное создание заказа",
			UserID: 4,
			Cart: &cartstorage.Cart{
				Items: map[int64]*cartstorage.CartItem{
					1: &cartstorage.CartItem{
						Quantity: 1,
					},
				},
			},
			OrderID: 1,
		},
	}

	sp := suite.NewSuiteProvider(t)

	cartService := cartservice.NewService(
		sp.GetProductService(),
		sp.GetCartStorage(),
		sp.GetLomsService(),
	)

	for _, test := range tests {

		sp.GetCartStorageMock().GetItemsByUserIDMock.
			When(ctx, test.UserID).
			Then(test.Cart, test.GetItemsByUserIDError)

		var orderItems lomsservice.OrderItems
		var orderCreateReq *order.OrderCreateRequest
		var orderCreateResp *order.OrderCreateResponse
		if test.Cart != nil {
			orderItems = cartservice.ToOrderItems(test.Cart.Items)
			orderCreateReq = lomsservice.ToOrderCreateRequest(test.UserID, orderItems)
			orderCreateResp = &order.OrderCreateResponse{OrderId: test.OrderID}
		}

		sp.GetLomsServiceOrderMock().CreateMock.
			When(ctx, orderCreateReq).
			Then(orderCreateResp, test.CreateOrderError)

		sp.GetCartStorageMock().DeleteItemsByUserIDMock.
			When(ctx, test.UserID).
			Then(test.DeleteItemsByUserIDerror)

		t.Run(test.Name, func(t *testing.T) {
			orderID, err := cartService.Checkout(ctx, test.UserID)
			if test.Error == nil {
				require.Nil(t, err, "Ошибки быть не должно")
				require.Equal(t, test.OrderID, orderID, "Не совпал номер заказа")

			} else {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не совпала ошибка")
			}
		})

	}

}
