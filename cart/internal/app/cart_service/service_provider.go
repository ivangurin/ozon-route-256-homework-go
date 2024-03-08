package cartservice

import (
	productservice "route256.ozon.ru/project/cart/internal/pkg/client/product_service"
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

type tServiceProvider struct {
	productService productservice.IClient
	cartStorage    cartstorage.IStorage
	cartService    cartservice.IService
}

var serviceProvider *tServiceProvider

func GetServiceProvider() *tServiceProvider {
	if serviceProvider == nil {
		serviceProvider = &tServiceProvider{}
	}
	return serviceProvider
}

func (sp *tServiceProvider) GetProductService() productservice.IClient {
	if sp.productService == nil {
		sp.productService = productservice.NewClient()
	}
	return sp.productService
}

func (sp *tServiceProvider) GetCartStorage() cartstorage.IStorage {
	if sp.cartStorage == nil {
		sp.cartStorage = cartstorage.NewCartStorage()
	}
	return sp.cartStorage
}

func (sp *tServiceProvider) GetCartService() cartservice.IService {
	if sp.cartService == nil {
		sp.cartService = cartservice.NewService(
			sp.GetProductService(),
			sp.GetCartStorage(),
		)
	}
	return sp.cartService
}
