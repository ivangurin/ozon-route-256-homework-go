package suite

import (
	"testing"

	"github.com/gojuno/minimock/v3"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	productservice_mocks "route256.ozon.ru/project/cart/internal/pkg/client/product_service/mocks"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
	cartstorage_mocks "route256.ozon.ru/project/cart/internal/repository/cart_storage/mocks"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
	cartservice_mocks "route256.ozon.ru/project/cart/internal/service/cart_service/mocks"
)

type suiteProvider struct {
	mc                 *minimock.Controller
	productServiceMock *productservice_mocks.ClientMockMock
	productService     productservice.IClient
	cartStorageMock    *cartstorage_mocks.StorageMockMock
	cartStorage        cartstorage.IStorage
	cartServiceMock    *cartservice_mocks.ServiceMockMock
	cartService        cartservice.IService
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

func (sp *suiteProvider) GetProductService() productservice.IClient {
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

func (sp *suiteProvider) GetCartStorege() cartstorage.IStorage {
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

func (sp *suiteProvider) GetCartService() cartservice.IService {
	if sp.cartService == nil {
		sp.cartService = cartservice.NewService(
			sp.GetProductService(),
			sp.GetCartStorege(),
		)
	}
	return sp.cartService
}
