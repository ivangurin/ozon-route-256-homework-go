package suite

import (
	kafka_mocks "route256.ozon.ru/project/loms/internal/pkg/kafka/mocks"
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
	kafkaProducer    *kafka_mocks.ProducerMock
	stockStorage     stockstorage.Repository
	stockStorageMock *stockstorage_mocks.RepositoryMock
	stockService     stockservce.Service
	stockServiceMock *stockservce_mocks.ServiceMock
	orderStorage     orderstorage.Repository
	orderStorageMock *orderstorage_mocks.RepositoryMock
	orderService     orderservice.Service
	orderServiceMock *orderservice_mocks.ServiceMock
}

func NewSuiteProvider() *suiteProvider {
	return &suiteProvider{}
}

func (sp *suiteProvider) GetKafkaProducer() *kafka_mocks.ProducerMock {
	if sp.kafkaProducer == nil {
		sp.kafkaProducer = &kafka_mocks.ProducerMock{}
	}
	return sp.kafkaProducer
}

func (sp *suiteProvider) GetStockStorageMock() *stockstorage_mocks.RepositoryMock {
	if sp.stockStorageMock == nil {
		sp.stockStorageMock = &stockstorage_mocks.RepositoryMock{}
	}
	return sp.stockStorageMock
}

func (sp *suiteProvider) GetStockStorage() stockstorage.Repository {
	if sp.stockStorage == nil {
		sp.stockStorage = sp.GetStockStorageMock()
	}
	return sp.stockStorage
}

func (sp *suiteProvider) GetOrderStorageMock() *orderstorage_mocks.RepositoryMock {
	if sp.orderStorageMock == nil {
		sp.orderStorageMock = &orderstorage_mocks.RepositoryMock{}
	}
	return sp.orderStorageMock
}

func (sp *suiteProvider) GetOrderStorage() orderstorage.Repository {
	if sp.orderStorage == nil {
		sp.orderStorage = sp.GetOrderStorageMock()
	}
	return sp.orderStorage
}

func (sp *suiteProvider) GetStockServiceMock() *stockservce_mocks.ServiceMock {
	if sp.stockServiceMock == nil {
		sp.stockServiceMock = &stockservce_mocks.ServiceMock{}
	}
	return sp.stockServiceMock
}

func (sp *suiteProvider) GetStockService() stockservce.Service {
	if sp.stockService == nil {
		sp.stockService = stockservce.NewService(
			sp.GetStockStorage(),
		)
	}
	return sp.stockService
}

func (sp *suiteProvider) GetOrderServiceMock() *orderservice_mocks.ServiceMock {
	if sp.orderServiceMock == nil {
		sp.orderServiceMock = &orderservice_mocks.ServiceMock{}
	}
	return sp.orderServiceMock
}

func (sp *suiteProvider) GetOrderService() orderservice.Service {
	if sp.orderService == nil {
		sp.orderService = orderservice.NewService(
			sp.GetStockStorage(),
			sp.GetOrderStorage(),
			sp.GetKafkaProducer(),
		)
	}
	return sp.orderService
}
