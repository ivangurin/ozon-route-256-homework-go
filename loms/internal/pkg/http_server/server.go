package httpserver

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"route256.ozon.ru/project/loms/internal/pkg/http_server/middleware"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

type API interface {
	RegisterHttpHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

type Server interface {
	Start() error
	Stop() error
	RegisterAPI(api []API) error
}

type server struct {
	ctx        context.Context
	mux        *http.ServeMux
	gwmux      *runtime.ServeMux
	conn       *grpc.ClientConn
	httpServer *http.Server
}

func NewServer(ctx context.Context, httpPort, grpcPort string) (Server, error) {

	var err error

	s := &server{
		ctx:   ctx,
		mux:   http.NewServeMux(),
		gwmux: runtime.NewServeMux(),
	}

	s.conn, err = grpc.Dial(
		fmt.Sprintf(":%s", grpcPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to deal: %w", err)
	}

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%s", httpPort),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		Handler:           middleware.WithHTTPLoggingMiddleware(s.mux),
	}

	// swagger
	s.mux.HandleFunc("/swagger.json", s.handleSwagger)

	// swagger-ui
	fs := http.FileServer(http.Dir("pkg/swagger-ui"))
	s.mux.Handle("/docs/", http.StripPrefix("/docs/", fs))

	// metrics
	s.mux.Handle("/metrics", promhttp.Handler())

	// grpc-gateway
	s.mux.Handle("/", s.gwmux)

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
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		logger.Errorf(ctx, "failed to stop http server: +v", err)
		return fmt.Errorf("failed to stop http server: %w", err)
	}
	logger.Info(ctx, "http server is stopped successfully")
	return nil
}

func (s *server) RegisterAPI(api []API) error {
	for _, singleAPI := range api {
		err := singleAPI.RegisterHttpHandler(s.ctx, s.gwmux, s.conn)
		if err != nil {
			return nil
		}
	}
	return nil
}

func (s *server) handleSwagger(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("pkg/swagger/swagger.json")
	if err != nil {
		http.Error(w, "swagger not found", http.StatusNotFound)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	_, err = io.Copy(w, reader)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}
