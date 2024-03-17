package cartservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestDeleteItem(t *testing.T) {

	type test struct {
		Name   string
		UserID int64
		SkuID  int64
		Error  error
	}

	var errCartNotFount = fmt.Errorf("cart not found")
	var errCartItemNotFound = fmt.Errorf("cart item not found")

	tests := []*test{
		{
			Name:   "Корзина не существует",
			UserID: 1,
			SkuID:  1,
			Error:  errCartNotFount,
		},
		{
			Name:   "Позция не существует",
			UserID: 2,
			SkuID:  1,
			Error:  errCartItemNotFound,
		},
		{
			Name:   "Корзина и позиция существуют",
			UserID: 3,
			SkuID:  1,
		},
	}

	ctx := context.Background()

	sp := suite.NewSuiteProvider(t)

	cartService := cartservice.NewService(
		sp.GetProductService(),
		sp.GetCartStorege(),
	)

	for _, test := range tests {

		sp.GetCartStoregeMock().DeleteItemMock.
			When(ctx, test.UserID, test.SkuID).
			Then(test.Error)

		t.Run(test.Name, func(t *testing.T) {

			err := cartService.DeleteItem(ctx, test.UserID, test.SkuID)
			require.ErrorIs(t, err, test.Error, "Должна быть ошибка")

		})
	}

}
