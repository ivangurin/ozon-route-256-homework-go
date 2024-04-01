package suite

import (
	"context"
	"testing"

	"github.com/gojuno/minimock/v3"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	orderstorage_mocks "route256.ozon.ru/project/loms/internal/repository/order_storage/mocks"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
	stockstorage_mocks "route256.ozon.ru/project/loms/internal/repository/stock_storage/mocks"
	orderservice "route256.ozon.ru/project/loms/internal/service/order_service"
	orderservice_mocks "route256.ozon.ru/project/loms/internal/service/order_service/mocks"
	stockservce "route256.ozon.ru/project/loms/internal/service/stock_service"
	stockservce_mocks "route256.ozon.ru/project/loms/internal/service/stock_service/mocks"
)

type suiteProvider struct {
	ctx              context.Context
	mc               *minimock.Controller
	stockStorage     stockstorage.Repository
	stockStorageMock *stockstorage_mocks.RepositoryMockMock
	stockService     stockservce.Service
	stockServiceMock *stockservce_mocks.ServiceMockMock
	orderStorage     orderstorage.Repository
	orderStorageMock *orderstorage_mocks.RepositoryMockMock
	orderService     orderservice.Service
	orderServiceMock *orderservice_mocks.ServiceMockMock
}

func NewSuiteProvider(t *testing.T, ctx context.Context) *suiteProvider {
	return &suiteProvider{
		ctx: ctx,
		mc:  minimock.NewController(t),
	}
}

func (sp *suiteProvider) GetStockStorageMock() *stockstorage_mocks.RepositoryMockMock {
	if sp.stockStorageMock == nil {
		sp.stockStorageMock = stockstorage_mocks.NewRepositoryMockMock(sp.mc)
	}
	return sp.stockStorageMock
}

func (sp *suiteProvider) GetStockStorage() stockstorage.Repository {
	if sp.stockStorage == nil {
		sp.stockStorage = sp.GetStockStorageMock()
	}
	return sp.stockStorage
}

func (sp *suiteProvider) GetOrderStorageMock() *orderstorage_mocks.RepositoryMockMock {
	if sp.orderStorageMock == nil {
		sp.orderStorageMock = orderstorage_mocks.NewRepositoryMockMock(sp.mc)
	}
	return sp.orderStorageMock
}

func (sp *suiteProvider) GetOrderStorage() orderstorage.Repository {
	if sp.orderStorage == nil {
		sp.orderStorage = sp.GetOrderStorageMock()
	}
	return sp.orderStorage
}

func (sp *suiteProvider) GetStockServiceMock() *stockservce_mocks.ServiceMockMock {
	if sp.stockServiceMock == nil {
		sp.stockServiceMock = stockservce_mocks.NewServiceMockMock(sp.mc)
	}
	return sp.stockServiceMock
}

func (sp *suiteProvider) GetStockService() stockservce.Service {
	if sp.stockService == nil {
		sp.stockService = stockservce.NewService(
			sp.ctx,
			sp.GetStockStorage(),
		)
	}
	return sp.stockService
}

func (sp *suiteProvider) GetOrderServiceMock() *orderservice_mocks.ServiceMockMock {
	if sp.orderServiceMock == nil {
		sp.orderServiceMock = orderservice_mocks.NewServiceMockMock(sp.mc)
	}
	return sp.orderServiceMock
}

func (sp *suiteProvider) GetOrderService() orderservice.Service {
	if sp.orderService == nil {
		sp.orderService = orderservice.NewService(
			sp.ctx,
			sp.GetStockStorage(),
			sp.GetOrderStorage(),
		)
	}
	return sp.orderService
}
