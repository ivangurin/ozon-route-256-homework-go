package serviceprovider

import (
	"context"

	"route256.ozon.ru/project/loms/internal/service/kafka_service"
	orderservice "route256.ozon.ru/project/loms/internal/service/order_service"
	stockservice "route256.ozon.ru/project/loms/internal/service/stock_service"
)

type services struct {
	orderService orderservice.Service
	stockService stockservice.Service
	kafkaService kafka_service.Service
}

func (sp *ServiceProvider) GetStockService(ctx context.Context) stockservice.Service {
	if sp.services.stockService == nil {
		sp.services.stockService = stockservice.NewService(
			sp.GetStockStorage(ctx),
		)
	}
	return sp.services.stockService
}

func (sp *ServiceProvider) GetOrderService(ctx context.Context) orderservice.Service {
	if sp.services.orderService == nil {
		sp.services.orderService = orderservice.NewService(
			sp.GetStockStorage(ctx),
			sp.GetOrderStorage(ctx),
			sp.GetSyncProducer(),
		)
	}
	return sp.services.orderService
}

func (sp *ServiceProvider) GetKafkaService(ctx context.Context) kafka_service.Service {
	if sp.services.kafkaService == nil {
		sp.services.kafkaService = kafka_service.NewService(
			sp.GetKafkaStorage(ctx),
			sp.GetSyncProducer(),
		)
	}
	return sp.services.kafkaService
}
