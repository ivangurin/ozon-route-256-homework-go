package cartservice_test

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/model"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestGetItemsByUserID(t *testing.T) {

	type test struct {
		Name   string
		UserID int64
		Cart   *cartservice.Cart
		Error  error
	}

	tests := []*test{
		{
			Name:   "Корзина не существует",
			UserID: 1,
			Error:  model.ErrNotFound,
		},
		{
			Name:   "Породукт не существует",
			UserID: 2,
			Error:  model.ErrNotFound,
		},
		{
			Name:   "Корзина существует",
			UserID: 3,
			Cart: &cartservice.Cart{
				Items: []*cartservice.CartItem{
					{
						SkuID:    1,
						Name:     "Product 1",
						Quantity: 1,
						Price:    100,
					},
					{
						SkuID:    2,
						Name:     "Product 2",
						Quantity: 2,
						Price:    200,
					},
					{
						SkuID:    3,
						Name:     "Product 3",
						Quantity: 3,
						Price:    300,
					},
				},
				TotalPrice: 1400,
			},
		},
	}

	ctx := context.Background()

	sp := suite.NewSuiteProvider(t)

	cartService := cartservice.NewService(
		sp.GetProductService(),
		sp.GetCartStorege(),
	)

	sp.GetCartStoregeMock().GetItemsByUserIDMock.
		When(ctx, 1).
		Then(nil, model.ErrNotFound)
	sp.GetCartStoregeMock().GetItemsByUserIDMock.
		When(ctx, 2).
		Then(&cartstorage.Cart{
			Items: cartstorage.CartItems{
				4: &cartstorage.CartItem{
					Quantity: 1,
				},
			},
		}, nil)
	sp.GetCartStoregeMock().GetItemsByUserIDMock.
		When(ctx, 3).
		Then(&cartstorage.Cart{
			Items: cartstorage.CartItems{
				1: &cartstorage.CartItem{
					Quantity: 1,
				},
				2: &cartstorage.CartItem{
					Quantity: 2,
				},
				3: &cartstorage.CartItem{
					Quantity: 3,
				},
			},
		}, nil)

	sp.GetProductServiceMock().GetProductWithRetriesMock.
		When(ctx, 1).
		Then(&productservice.GetProductResponse{
			Name:  "Product 1",
			Price: 100,
		}, nil)
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
	sp.GetProductServiceMock().GetProductWithRetriesMock.
		When(ctx, 4).
		Then(nil, model.ErrNotFound)

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {

			cart, err := cartService.GetItemsByUserID(ctx, test.UserID)
			require.ErrorIs(t, err, test.Error, "Должна быть ошибка")

			diff := deep.Equal(cart, test.Cart)
			if diff != nil {
				t.Errorf("Корзины должны совпадать: %+v", diff)
			}

		})
	}

}
