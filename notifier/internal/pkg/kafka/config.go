package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

func NewConfig(opts ...ConsumerGroupOption) *sarama.Config {
	config := sarama.NewConfig()

	config.Version = sarama.MaxVersion

	/*
		sarama.OffsetNewest - получаем только новые сообщений, те, которые уже были игнорируются
		sarama.OffsetOldest - читаем все с самого начала
	*/
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Используется, если ваш offset "уехал" далеко и нужно пропустить невалидные сдвиги
	config.Consumer.Group.ResetInvalidOffsets = true

	// Сердцебиение консьюмера
	config.Consumer.Group.Heartbeat.Interval = 3 * time.Second

	// Таймаут сессии
	config.Consumer.Group.Session.Timeout = 60 * time.Second

	// Таймаут ребалансировки
	config.Consumer.Group.Rebalance.Timeout = 60 * time.Second

	//
	config.Consumer.Return.Errors = true

	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second

	for _, opt := range opts {
		opt(config)
	}

	return config
}
