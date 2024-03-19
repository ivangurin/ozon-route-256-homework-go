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
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: middleware.WithHTTPLoggingMiddleware(s.mux),
	}

	s.mux.HandleFunc("/swagger.json", s.handleSwagger)

	fs := http.FileServer(http.Dir("pkg/swagger-ui"))
	s.mux.Handle("/docs/", http.StripPrefix("/docs/", fs))

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
	return s.httpServer.Shutdown(ctx)
}

func (s *server) RegisterAPI(api api) error {
	return api.RegisterHttpHandler(s.ctx, s.gwmux, s.conn)
}

func (s *server) handleSwagger(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("pkg/swagger/swagger.json")
	if err != nil {
		http.Error(w, "swagger not found", http.StatusNotFound)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	io.Copy(w, reader)
}
