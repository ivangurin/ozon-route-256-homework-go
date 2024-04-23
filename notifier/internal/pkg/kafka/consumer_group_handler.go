package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"go.opentelemetry.io/otel/trace"
	"route256.ozon.ru/project/notifier/internal/pkg/logger"
)

type Handler func(ctx context.Context, message *sarama.ConsumerMessage) (bool, error)

type consumerGroupHandler struct {
	ctx     context.Context
	handler Handler
}

func NewConsumerGroupHandler(
	ctx context.Context,
	handler Handler,
) *consumerGroupHandler {
	return &consumerGroupHandler{
		ctx:     ctx,
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
//
//nolint:gocognit
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}

			ctx := session.Context()
			var traceID string
			var spanID string

			for _, header := range message.Headers {
				switch string(header.Key) {
				case "x-trace-id":
					traceID = string(header.Value)
				case "x-span-id":
					spanID = string(header.Value)
				}
			}

			if traceID != "" {
				spanContext := trace.SpanContextConfig{
					TraceFlags: trace.FlagsSampled,
					Remote:     true,
				}
				spanContext.TraceID, _ = trace.TraceIDFromHex(traceID)
				spanContext.SpanID, _ = trace.SpanIDFromHex(spanID)
				ctx = trace.ContextWithSpanContext(ctx,
					trace.NewSpanContext(spanContext))
			}

			ok, err := h.handler(ctx, message)
			if err != nil {
				logger.Errorf(h.ctx, "failed to handle message: %v", err)
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
