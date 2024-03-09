package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"route256.ozon.ru/project/cart/internal/model"
)

type IServer interface {
	AddHandlers(handlers model.HttpApiHandlers)
	Start() error
	Stop() error
}

type server struct {
	server http.Server
}

func NewServer(port string) IServer {
	s := &server{
		server: http.Server{
			Addr: fmt.Sprintf(":%s", port),
		},
	}

	return s
}

func (s *server) AddHandlers(handlers model.HttpApiHandlers) {
	for _, handeler := range handlers {
		http.HandleFunc(handeler.Pattern, handeler.Handler)
	}
}

func (s *server) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to start http server: %w", err)
		}
	}

	return nil
}

func (s *server) Stop() error {
	return s.server.Shutdown(context.Background())
}
