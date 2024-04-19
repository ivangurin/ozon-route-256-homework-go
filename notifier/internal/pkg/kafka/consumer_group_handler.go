package kafka

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"go.opentelemetry.io/otel/trace"
	"route256.ozon.ru/project/notifier/internal/pkg/logger"
	"route256.ozon.ru/project/notifier/internal/pkg/tracer"
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
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return nil
			}

			ctx := session.Context()

			for _, header := range message.Headers {
				key := string(header.Key)

				if key == "x-trace-id" {
					traceID := string(header.Value)
					if traceID != "" {
						fmt.Println("traceID", traceID)
						traceIDHex, _ := trace.TraceIDFromHex(traceID)
						spanContext := trace.NewSpanContext(trace.SpanContextConfig{
							TraceID: traceIDHex,
						})
						ctx = trace.ContextWithSpanContext(ctx, spanContext)
					}
				}

			}

			var span trace.Span
			ctx, span = tracer.StartSpanFromContext(ctx, "kafkaConsumerGroup.HandleMessage")
			defer span.End()

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
