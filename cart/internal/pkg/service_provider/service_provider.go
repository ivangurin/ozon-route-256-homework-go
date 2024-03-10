package serviceprovider

import (
	"os"
	"syscall"

	cartapi "route256.ozon.ru/project/cart/internal/api/cart_api"
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	"route256.ozon.ru/project/cart/internal/pkg/closer"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

type ServiceProvider struct {
	closer         closer.ICloser
	cartAPI        cartapi.IAPI
	productService productservice.IClient
	cartStorage    cartstorage.IStorage
	cartService    cartservice.IService
}

var serviceProvider *ServiceProvider

func GetServiceProvider() *ServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &ServiceProvider{}
	}
	return serviceProvider
}

func (sp *ServiceProvider) GetCloser() closer.ICloser {
	if sp.closer == nil {
		sp.closer = closer.NewCloser(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	}
	return sp.closer
}

func (sp *ServiceProvider) GetCartAPI() cartapi.IAPI {
	if sp.cartAPI == nil {
		sp.cartAPI = cartapi.NewAPI(sp.GetCartService())
	}
	return sp.cartAPI
}

func (sp *ServiceProvider) GetProductService() productservice.IClient {
	if sp.productService == nil {
		sp.productService = productservice.NewClient()
	}
	return sp.productService
}

func (sp *ServiceProvider) GetCartStorage() cartstorage.IStorage {
	if sp.cartStorage == nil {
		sp.cartStorage = cartstorage.NewCartStorage()
	}
	return sp.cartStorage
}

func (sp *ServiceProvider) GetCartService() cartservice.IService {
	if sp.cartService == nil {
		sp.cartService = cartservice.NewService(
			sp.GetProductService(),
			sp.GetCartStorage(),
		)
	}
	return sp.cartService
}
