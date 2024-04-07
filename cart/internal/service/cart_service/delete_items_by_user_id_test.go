package cartservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestDeleteItemsByUserID(t *testing.T) {

	type test struct {
		Name   string
		UserID int64
		Error  error
	}

	var errCartNotFount = fmt.Errorf("cart not found")

	tests := []*test{
		{
			Name:   "Корзина не существует",
			UserID: 1,
			Error:  errCartNotFount,
		},
		{
			Name:   "Корзина существуют",
			UserID: 2,
		},
	}

	t.Parallel()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()

			cartService := cartservice.NewService(
				sp.GetProductService(),
				sp.GetCartStorage(),
				sp.GetLomsService(),
			)

			sp.GetCartStorageMock().EXPECT().
				DeleteItemsByUserID(mock.Anything, test.UserID).
				Return(test.Error)

			err := cartService.DeleteItemsByUserID(context.Background(), test.UserID)
			require.ErrorIs(t, err, test.Error, "Должна быть ошибка")
		})
	}
}
