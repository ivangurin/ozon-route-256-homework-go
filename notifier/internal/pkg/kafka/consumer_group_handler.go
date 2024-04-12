package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"route256.ozon.ru/project/notifier/internal/pkg/logger"
)

type Handler func(ctx context.Context, message *sarama.ConsumerMessage) (bool, error)

type consumerGroupHandler struct {
	handler Handler
}

func NewConsumerGroupHandler(
	handler Handler,
) *consumerGroupHandler {
	return &consumerGroupHandler{
		handler: handler,
	}
}

// Setup Начинаем новую сессию, до ConsumeClaim
func (h *consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup завершает сессию, после того, как все ConsumeClaim завершатся
func (h *consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim читаем до тех пор пока сессия не завершилась
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}

			ok, err := h.handler(session.Context(), message)
			if err != nil {
				logger.Errorf("failed to handle message: %v", err)
				continue
			}

			if !ok {
				continue
			}

			// mark message as successfully handled and ready to commit offset
			session.MarkMessage(message, "")

			// commit offset manually right now
			// works when autocommit disabled
			// session.Commit()

		case <-session.Context().Done():
			return nil
		}
	}
}
