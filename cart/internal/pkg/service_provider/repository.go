package serviceprovider

import (
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
	cartservice "route256.ozon.ru/project/cart/internal/service/cart_service"
)

type repositories struct {
	cartStorage cartstorage.Storage
	cartService cartservice.Service
}

func (sp *ServiceProvider) GetCartStorage() cartstorage.Storage {
	if sp.repositories.cartStorage == nil {
		sp.repositories.cartStorage = cartstorage.NewCartStorage()
	}
	return sp.repositories.cartStorage
}
