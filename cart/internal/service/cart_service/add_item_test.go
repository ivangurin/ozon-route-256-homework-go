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
		Name         string
		UserID       int64
		SkuID        int64
		Qunatity     uint32
		Product      *productservice.GetProductResponse
		ProductError error
		Error        error
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
			Product: &productservice.GetProductResponse{
				Name:  "Product 2",
				Price: 200,
			},
			Error: someError,
		},
		{
			Name: "Продукт успешно добавлен",
			Product: &productservice.GetProductResponse{
				Name:  "Product 3",
				Price: 300,
			},
			SkuID: 3,
		},
	}

	ctx := context.Background()

	sp := suite.NewSuiteProvider(t)

	cartService := cartservice.NewService(
		sp.GetProductService(),
		sp.GetCartStorege(),
	)

	for _, test := range tests {

		sp.GetProductServiceMock().GetProductWithRetriesMock.
			When(ctx, test.SkuID).
			Then(test.Product, test.Error)

		sp.GetCartStoregeMock().AddItemMock.
			When(ctx, test.UserID, test.SkuID, uint16(test.Qunatity)).
			Then(test.Error)

		t.Run(test.Name, func(t *testing.T) {

			err := cartService.AddItem(ctx, test.UserID, test.SkuID, uint16(test.Qunatity))
			require.ErrorIs(t, err, test.Error, "Должна быть ошибка NotFound")

		})
	}

}
