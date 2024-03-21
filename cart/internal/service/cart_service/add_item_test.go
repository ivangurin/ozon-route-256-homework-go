package cartservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestAddItem(t *testing.T) {

	type test struct {
		Name             string
		UserID           int64
		SkuID            int64
		Qunatity         uint32
		ProductInfo      *productservice.GetProductResponse
		ProductInfoError error
		StockInfoReq     *stock.StockInfoRequest
		StockInfoResp    *stock.StockInfoResponse
		StockInfoError   error
		StorageAddError  error
		Error            error
	}

	var (
		err1 = fmt.Errorf("some error 1")
		err2 = fmt.Errorf("some error 2")
		err3 = fmt.Errorf("some error 3")
		err4 = fmt.Errorf("some error 4")
	)

	ctx := context.Background()

	tests := []*test{
		{
			Name:             "Продукт не существует",
			SkuID:            1,
			ProductInfoError: err1,
			Error:            err1,
		},
		{
			Name:  "Ошибка при получении резева",
			SkuID: 2,
			ProductInfo: &productservice.GetProductResponse{
				Name:  "Product 2",
				Price: 200,
			},
			StockInfoReq: &stock.StockInfoRequest{
				Sku: 2,
			},
			StockInfoError: err2,
			Error:          err2,
		},
		{
			Name:  "В резерве меньше чем нужно",
			SkuID: 3,
			ProductInfo: &productservice.GetProductResponse{
				Name:  "Product 2",
				Price: 200,
			},
			StockInfoReq: &stock.StockInfoRequest{
				Sku: 3,
			},
			StockInfoResp: &stock.StockInfoResponse{
				Count: 100,
			},
			StockInfoError: err3,
			Error:          err3,
		},
		{
			Name:  "Ошибка при добавлении в сторадж",
			SkuID: 4,
			ProductInfo: &productservice.GetProductResponse{
				Name:  "Product 2",
				Price: 200,
			},
			StockInfoReq: &stock.StockInfoRequest{
				Sku: 4,
			},
			StockInfoResp: &stock.StockInfoResponse{
				Count: 1000,
			},
			StorageAddError: err4,
			Error:           err4,
		},
		{
			Name:  "Продукт успешно добавлен",
			SkuID: 5,
			ProductInfo: &productservice.GetProductResponse{
				Name:  "Product 3",
				Price: 300,
			},
			StockInfoReq: &stock.StockInfoRequest{
				Sku: 5,
			},
			StockInfoResp: &stock.StockInfoResponse{
				Count: 1000,
			},
		},
	}

	sp := suite.NewSuiteProvider(t)

	cartService := cartservice.NewService(
		sp.GetProductService(),
		sp.GetCartStorege(),
		sp.GetLomsService(),
	)

	for _, test := range tests {

		sp.GetProductServiceMock().GetProductWithRetriesMock.
			When(ctx, test.SkuID).
			Then(test.ProductInfo, test.ProductInfoError)

		sp.GetLomsServiceStockMock().InfoMock.
			When(ctx, test.StockInfoReq).
			Then(test.StockInfoResp, test.StockInfoError)

		sp.GetCartStoregeMock().AddItemMock.
			When(ctx, test.UserID, test.SkuID, uint16(test.Qunatity)).
			Then(test.StorageAddError)

		t.Run(test.Name, func(t *testing.T) {
			err := cartService.AddItem(ctx, test.UserID, test.SkuID, uint16(test.Qunatity))
			if test.Error != nil {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не совпала ошибка")
			} else {
				require.Nil(t, err, "Ошибки быть не должно")
			}
		})
	}

}
