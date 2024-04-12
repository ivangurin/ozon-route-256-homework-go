package serviceprovider

import (
	"os"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"route256.ozon.ru/project/loms/internal/config"
	"route256.ozon.ru/project/loms/internal/pkg/closer"
	"route256.ozon.ru/project/loms/internal/pkg/kafka"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

type ServiceProvider struct {
	closer       closer.ICloser
	syncProducer kafka.Producer

	api          api
	clients      clients
	services     services
	repositories repositories
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

func (sp *ServiceProvider) GetSyncProducer() kafka.Producer {
	if sp.syncProducer == nil {
		var err error
		sp.syncProducer, err = kafka.NewSyncProducer(
			config.KafkaAddr,
			kafka.WithIdempotent(),
			kafka.WithRequiredAcks(sarama.WaitForAll),
			kafka.WithMaxOpenRequests(1),
			kafka.WithMaxRetries(5),
			kafka.WithRetryBackoff(10*time.Millisecond),
			// kafka.WithProducerPartitioner(sarama.NewRoundRobinPartitioner),
		)
		if err != nil {
			logger.Fatalf("failed to create kafka producer: %v", err)
		}
		sp.GetCloser().Add(sp.syncProducer.Close)
	}
	return sp.syncProducer
}
