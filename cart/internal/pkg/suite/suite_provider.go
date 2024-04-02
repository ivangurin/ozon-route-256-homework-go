package suite

import (
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
	productServiceMock   *productservice_mocks.ClientMock
	productService       productservice.Client
	cartStorageMock      *cartstorage_mocks.StorageMock
	cartStorage          cartstorage.Storage
	cartServiceMock      *cartservice_mocks.ServiceMock
	cartService          cartservice.Service
	lomsServiceStockMock *lomsservice_mocks.StockClientMock
	lomsServiceOrderMock *lomsservice_mocks.OrderClientMock
	lomsService          *lomsservice.Client
}

func NewSuiteProvider() *suiteProvider {
	return &suiteProvider{}
}

func (sp *suiteProvider) GetProductServiceMock() *productservice_mocks.ClientMock {
	if sp.productServiceMock == nil {
		sp.productServiceMock = &productservice_mocks.ClientMock{}
	}
	return sp.productServiceMock
}

func (sp *suiteProvider) GetProductService() productservice.Client {
	if sp.productService == nil {
		sp.productService = sp.GetProductServiceMock()
	}
	return sp.productService
}

func (sp *suiteProvider) GetCartStorageMock() *cartstorage_mocks.StorageMock {
	if sp.cartStorageMock == nil {
		sp.cartStorageMock = &cartstorage_mocks.StorageMock{}
	}
	return sp.cartStorageMock
}

func (sp *suiteProvider) GetCartStorage() cartstorage.Storage {
	if sp.cartStorage == nil {
		sp.cartStorage = sp.GetCartStorageMock()
	}
	return sp.cartStorage
}

func (sp *suiteProvider) GetCartServiceMock() *cartservice_mocks.ServiceMock {
	if sp.cartServiceMock == nil {
		sp.cartServiceMock = &cartservice_mocks.ServiceMock{}
	}
	return sp.cartServiceMock
}

func (sp *suiteProvider) GetCartService() cartservice.Service {
	if sp.cartService == nil {
		sp.cartService = cartservice.NewService(
			sp.GetProductService(),
			sp.GetCartStorage(),
			sp.GetLomsService(),
		)
	}
	return sp.cartService
}

func (sp *suiteProvider) GetLomsServiceStockMock() *lomsservice_mocks.StockClientMock {
	if sp.lomsServiceStockMock == nil {
		sp.lomsServiceStockMock = &lomsservice_mocks.StockClientMock{}
	}
	return sp.lomsServiceStockMock
}

func (sp *suiteProvider) GetLomsServiceOrderMock() *lomsservice_mocks.OrderClientMock {
	if sp.lomsServiceOrderMock == nil {
		sp.lomsServiceOrderMock = &lomsservice_mocks.OrderClientMock{}
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
