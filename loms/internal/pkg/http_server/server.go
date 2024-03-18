package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"route256.ozon.ru/project/loms/internal/pkg/middleware"
)

type api interface {
	RegisterHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type Server interface {
	Start() error
	Stop() error
	RegisterAPI(api api) error
}

type server struct {
	ctx        context.Context
	gwmux      *runtime.ServeMux
	conn       *grpc.ClientConn
	httpServer *http.Server
}

func NewServer(ctx context.Context, httpPort, grpcPort string) (Server, error) {
	s := &server{
		ctx:   ctx,
		gwmux: runtime.NewServeMux(),
	}

	var err error
	s.conn, err = grpc.Dial(
		fmt.Sprintf(":%s", grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to deal: %w", err)
	}

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: middleware.WithHTTPLoggingMiddleware(s.gwmux),
	}

	return s, nil
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
	ctx, cancel := context.WithTimeout(s.ctx, 3*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

func (s *server) RegisterAPI(api api) error {
	return api.RegisterHttpHandler(s.ctx, s.gwmux, s.conn)
}
