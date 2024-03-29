package cartservice_test

import (
	"context"
	"fmt"
	"testing"

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

	ctx := context.Background()

	sp := suite.NewSuiteProvider(t)
	cartService := cartservice.NewService(
		sp.GetProductService(),
		sp.GetCartStorage(),
		sp.GetLomsService(),
	)

	for _, test := range tests {

		sp.GetCartStorageMock().DeleteItemsByUserIDMock.
			When(ctx, test.UserID).
			Then(test.Error)

		t.Run(test.Name, func(t *testing.T) {

			err := cartService.DeleteItemsByUserID(ctx, test.UserID)
			require.ErrorIs(t, err, test.Error, "Должна быть ошибка")

		})
	}

}
