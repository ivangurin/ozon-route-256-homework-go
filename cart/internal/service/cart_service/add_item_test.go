package cartservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/model"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestAddItem(t *testing.T) {

	type test struct {
		Name  string
		SkuID int64
		Error error
	}

	var someError = fmt.Errorf("some error")

	tests := []*test{
		{
			Name:  "Продукт не существует",
			SkuID: 1,
			Error: model.ErrNotFound,
		},
		{
			Name:  "Ошика при добавлении в сторадж",
			SkuID: 2,
			Error: someError,
		},
		{
			Name:  "Продукт успешно добавлен",
			SkuID: 3,
		},
	}

	ctx := context.Background()

	sp := suite.NewSuiteProvider(t)

	cartService := cartservice.NewService(
		sp.GetProductService(),
		sp.GetCartStorege(),
	)

	sp.GetProductServiceMock().GetProductWithRetriesMock.
		When(ctx, 1).
		Then(nil, model.ErrNotFound)

	sp.GetProductServiceMock().GetProductWithRetriesMock.
		When(ctx, 2).
		Then(&productservice.GetProductResponse{
			Name:  "Product 2",
			Price: 200,
		}, nil)

	sp.GetProductServiceMock().GetProductWithRetriesMock.
		When(ctx, 3).
		Then(&productservice.GetProductResponse{
			Name:  "Product 3",
			Price: 300,
		}, nil)

	sp.GetCartStoregeMock().AddItemMock.
		When(ctx, 0, 2, 0).
		Then(someError)

	sp.GetCartStoregeMock().AddItemMock.
		When(ctx, 0, 3, 0).
		Then(nil)

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {

			err := cartService.AddItem(ctx, 0, test.SkuID, 0)
			require.ErrorIs(t, err, test.Error, "Должна быть ошибка NotFound")

		})
	}

}
