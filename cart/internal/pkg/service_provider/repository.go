package serviceprovider

import (
	cartstorage "route256.ozon.ru/project/cart/internal/repository/cart_storage"
)

type repositories struct {
	cartStorage cartstorage.Storage
}

func (sp *ServiceProvider) GetCartStorage() cartstorage.Storage {
	if sp.repositories.cartStorage == nil {
		sp.repositories.cartStorage = cartstorage.NewCartStorage()
	}
	return sp.repositories.cartStorage
}
