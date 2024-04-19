package serviceprovider

import (
	"context"

	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/db"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
	"route256.ozon.ru/project/loms/internal/repository/kafka_storage"
	orderstorage "route256.ozon.ru/project/loms/internal/repository/order_storage"
	stockstorage "route256.ozon.ru/project/loms/internal/repository/stock_storage"
)

type repositories struct {
	dbClient     db.Client
	orderStorage orderstorage.Repository
	stockStorage stockstorage.Repository
	kafkaStorage kafka_storage.Repository
}

func (sp *ServiceProvider) GetDBClient(ctx context.Context) db.Client {
	if sp.repositories.dbClient == nil {
		dbc, err := db.NewClient(ctx, config.MasterDBUrl, config.SyncDBUrl)
		if err != nil {
			logger.Fatalf("failed to create db client: %v", err)
		}
		sp.repositories.dbClient = dbc
		sp.GetCloser().Add(dbc.Close)
	}
	return sp.repositories.dbClient
}

func (sp *ServiceProvider) GetStockStorage(ctx context.Context) stockstorage.Repository {
	if sp.repositories.stockStorage == nil {
		sp.repositories.stockStorage = stockstorage.NewRepository(
			ctx,
			sp.GetDBClient(ctx),
		)
	}
	return sp.repositories.stockStorage
}

func (sp *ServiceProvider) GetOrderStorage(ctx context.Context) orderstorage.Repository {
	if sp.repositories.orderStorage == nil {
		sp.repositories.orderStorage = orderstorage.NewRepository(
			ctx,
			sp.GetDBClient(ctx),
		)
	}
	return sp.repositories.orderStorage
}

func (sp *ServiceProvider) GetKafkaStorage(ctx context.Context) kafka_storage.Repository {
	if sp.repositories.kafkaStorage == nil {
		sp.repositories.kafkaStorage = kafka_storage.NewRepository(
			ctx,
			sp.GetDBClient(ctx),
		)
	}
	return sp.repositories.kafkaStorage
}
