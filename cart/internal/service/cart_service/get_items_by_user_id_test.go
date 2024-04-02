package cartservice_test

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/model"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestGetItemsByUserID(t *testing.T) {

	type testProduct struct {
		ID    int64
		Error error
		Data  *productservice.GetProductResponse
	}

	type test struct {
		Name        string
		UserID      int64
		Products    []*testProduct
		CartStorage *cartstorage.Cart
		CartService *cartservice.Cart
		Error       error
	}

	tests := []*test{
		{
			Name:   "Корзина не существует",
			UserID: 1,
			Error:  model.ErrNotFound,
		},
		{
			Name:   "Продукт не существует",
			UserID: 2,
			Error:  model.ErrNotFound,
		},
		{
			Name:   "Корзина существует",
			UserID: 3,
			Products: []*testProduct{
				{
					ID: 1,
					Data: &productservice.GetProductResponse{
						Name:  "Product 1",
						Price: 100,
					},
				},
				{
					ID: 2,
					Data: &productservice.GetProductResponse{
						Name:  "Product 2",
						Price: 200,
					},
				},
				{
					ID: 3,
					Data: &productservice.GetProductResponse{
						Name:  "Product 3",
						Price: 300,
					},
				},
			},
			CartStorage: &cartstorage.Cart{
				Items: map[int64]*cartstorage.CartItem{
					1: {Quantity: 1},
					2: {Quantity: 2},
					3: {Quantity: 3},
				},
			},
			CartService: &cartservice.Cart{
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
				GetItemsByUserID(mock.Anything, test.UserID).
				Return(test.CartStorage, test.Error)

			for _, testProduct := range test.Products {
				sp.GetProductServiceMock().EXPECT().
					GetProductWithRetries(mock.Anything, testProduct.ID).
					Return(testProduct.Data, testProduct.Error)
			}

			cart, err := cartService.GetItemsByUserID(context.Background(), test.UserID)
			if test.Error != nil {
				require.ErrorIs(t, err, test.Error, "Должна быть ошибка")
			} else {
				require.Nil(t, err, "Ошибки быть не должно")
			}

			diff := deep.Equal(cart, test.CartService)
			if diff != nil {
				t.Errorf("Корзины должны совпадать: %+v", diff)
			}
		})
	}
}
