package suite

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	lomsservice "route256.ozon.ru/project/cart/internal/pkg/client/loms_service"
	lomsservice_mocks "route256.ozon.ru/project/cart/internal/pkg/client/loms_service/mocks"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	productservice_mocks "route256.ozon.ru/project/cart/internal/pkg/client/product_service/mocks"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
	cartstorage_mocks "route256.ozon.ru/project/cart/internal/repository/cart_storage/mocks"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
	cartservice_mocks "route256.ozon.ru/project/cart/internal/service/cart_service/mocks"
)

type suiteProvider struct {
	mc                   *minimock.Controller
	productServiceMock   *productservice_mocks.ClientMockMock
	productService       productservice.Client
	cartStorageMock      *cartstorage_mocks.StorageMockMock
	cartStorage          cartstorage.Storage
	cartServiceMock      *cartservice_mocks.ServiceMockMock
	cartService          cartservice.Service
	lomsServiceStockMock *lomsservice_mocks.StockClientMockMock
	lomsServiceOrderMock *lomsservice_mocks.OrderClientMockMock
	lomsService          *lomsservice.Client
}

func NewSuiteProvider(t *testing.T) *suiteProvider {
	return &suiteProvider{
		mc: minimock.NewController(t),
	}
}

func (sp *suiteProvider) GetProductServiceMock() *productservice_mocks.ClientMockMock {
	if sp.productServiceMock == nil {
		sp.productServiceMock = productservice_mocks.NewClientMockMock(sp.mc)
	}
	return sp.productServiceMock
}

func (sp *suiteProvider) GetProductService() productservice.Client {
	if sp.productService == nil {
		sp.productService = sp.GetProductServiceMock()
	}
	return sp.productService
}

func (sp *suiteProvider) GetCartStoregeMock() *cartstorage_mocks.StorageMockMock {
	if sp.cartStorageMock == nil {
		sp.cartStorageMock = cartstorage_mocks.NewStorageMockMock(sp.mc)
	}
	return sp.cartStorageMock
}

func (sp *suiteProvider) GetCartStorege() cartstorage.Storage {
	if sp.cartStorage == nil {
		sp.cartStorage = sp.GetCartStoregeMock()
	}
	return sp.cartStorage
}

func (sp *suiteProvider) GetCartServiceMock() *cartservice_mocks.ServiceMockMock {
	if sp.cartServiceMock == nil {
		sp.cartServiceMock = cartservice_mocks.NewServiceMockMock(sp.mc)
	}
	return sp.cartServiceMock
}

func (sp *suiteProvider) GetCartService() cartservice.Service {
	if sp.cartService == nil {
		sp.cartService = cartservice.NewService(
			sp.GetProductService(),
			sp.GetCartStorege(),
			sp.GetLomsService(),
		)
	}
	return sp.cartService
}

func (sp *suiteProvider) GetLomsServiceStockMock() *lomsservice_mocks.StockClientMockMock {
	if sp.lomsServiceStockMock == nil {
		sp.lomsServiceStockMock = lomsservice_mocks.NewStockClientMockMock(sp.mc)
	}
	return sp.lomsServiceStockMock
}

func (sp *suiteProvider) GetLomsServiceOrderMock() *lomsservice_mocks.OrderClientMockMock {
	if sp.lomsServiceOrderMock == nil {
		sp.lomsServiceOrderMock = lomsservice_mocks.NewOrderClientMockMock(sp.mc)
	}
	return sp.lomsServiceOrderMock
}

func (sp *suiteProvider) GetLomsService() *lomsservice.Client {
	if sp.lomsService == nil {
		sp.lomsService = &lomsservice.Client{
			StockAPI: sp.GetLomsServiceStockMock(),
			OrderAPI: sp.GetLomsServiceOrderMock(),
		}
	}
	return sp.lomsService
}
