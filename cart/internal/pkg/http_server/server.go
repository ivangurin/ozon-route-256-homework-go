package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/internal/pkg/http_server/middleware"
	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

type IServer interface {
	AddHandlers(handlers model.HttpApiHandlers)
	Start() error
	Stop() error
}

type server struct {
	ctx        context.Context
	mux        *http.ServeMux
	httpServer *http.Server
}

func NewServer(ctx context.Context, port string) IServer {
	s := &server{
		ctx: ctx,
		mux: http.NewServeMux(),
	}

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		Handler:           s.mux,
	}

	// metrics
	s.mux.Handle("/metrics", promhttp.Handler())

	return s
}

func (s *server) AddHandlers(handlers model.HttpApiHandlers) {
	for _, handler := range handlers {
		s.mux.HandleFunc(
			handler.Pattern,
			middleware.WithLogger(
				middleware.WithMetrics(
					handler.Handler)))
	}
}

func (s *server) Start() error {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to start http server: %w", err)
		}
	}

	return nil
}

func (s *server) Stop() error {
	err := s.httpServer.Shutdown(s.ctx)
	if err != nil {
		logger.Errorf(s.ctx, "failed to stop http server: %v", err)
		return fmt.Errorf("failed to stop http server: %w", err)
	}
	logger.Info(s.ctx, "http server is stopped successfully")
	return nil
}
