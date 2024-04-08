package notifierservice

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/notifier/internal/config"
	"route256.ozon.ru/project/notifier/internal/consumers/order_status_changed_consumer"
	"route256.ozon.ru/project/notifier/internal/pkg/kafka"
	"route256.ozon.ru/project/notifier/internal/pkg/logger"
	serviceprovider "route256.ozon.ru/project/notifier/internal/pkg/service_provider"
)

type App interface {
	Run() error
}

type app struct {
	ctx context.Context
	sp  *serviceprovider.ServiceProvider
}

func NewApp(ctx context.Context) App {
	ctx, cancel := context.WithCancel(ctx)

	sp := serviceprovider.GetServiceProvider()
	sp.GetCloser().Add(func() error {
		cancel()
		return nil
	})

	return &app{
		ctx: ctx,
		sp:  sp,
	}
}

func (a *app) Run() error {
	var err error

	logger.Info("app is starting...")
	defer logger.Info("app finished")

	closer := a.sp.GetCloser()
	defer closer.Wait()

	consumer := order_status_changed_consumer.NewConsumer(
		a.sp.GetNotifierService(),
	)

	consumerGroup, err := kafka.NewConsumerGroup(
		a.ctx,
		[]string{config.KafkaAddr},
		config.KafkaConsumerGroupID,
		[]string{config.KafkaOrderEventsTopic},
		consumer.Handle,
	)
	if err != nil {
		logger.Errorf("failed to create consumer group: %v", err)
		closer.CloseAll()
		return fmt.Errorf("failed to create consumer group: %w", err)
	}
	closer.Add(consumerGroup.Close)

	go func() {
		logger.Info("consumer group is starting...")
		err := consumerGroup.Run()
		if err != nil {
			logger.Errorf("failed to start consumer group: %v", err)
			closer.CloseAll()
			return
		}
		logger.Info("consumer group started successfully")
	}()

	return err
}
