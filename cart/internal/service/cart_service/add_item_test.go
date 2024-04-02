package cartservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
	"route256.ozon.ru/project/cart/internal/pb/api/stock/v1"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	"route256.ozon.ru/project/cart/internal/pkg/suite"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

func TestTest(t *testing.T) {
	defer goleak.VerifyNone(t, goleak.IgnoreTopFunction("github.com/golang/glog.(*loggingT).flushDaemon"))

}

func TestAddItem(t *testing.T) {

	type test struct {
		Name             string
		UserID           int64
		SkuID            int64
		Quantity         uint32
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

	tests := []*test{
		{
			Name:             "Продукт не существует",
			SkuID:            1,
			ProductInfoError: err1,
			Error:            err1,
		},
		{
			Name:  "Ошибка при получении резерва",
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

	t.Parallel()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			sp := suite.NewSuiteProvider()

			cartService := cartservice.NewService(
				sp.GetProductService(),
				sp.GetCartStorage(),
				sp.GetLomsService(),
			)

			sp.GetProductServiceMock().EXPECT().
				GetProductWithRetries(mock.Anything, test.SkuID).
				Return(test.ProductInfo, test.ProductInfoError)

			sp.GetLomsServiceStockMock().EXPECT().
				Info(mock.Anything, test.StockInfoReq).
				Return(test.StockInfoResp, test.StockInfoError)

			sp.GetCartStorageMock().EXPECT().
				AddItem(mock.Anything, test.UserID, test.SkuID, uint16(test.Quantity)).
				Return(test.StorageAddError)

			err := cartService.AddItem(context.Background(), test.UserID, test.SkuID, uint16(test.Quantity))
			if test.Error != nil {
				require.NotNil(t, err, "Должна быть ошибка")
				require.ErrorIs(t, err, test.Error, "Не совпала ошибка")
			} else {
				require.Nil(t, err, "Ошибки быть не должно")
			}
		})
	}
}
