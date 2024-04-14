package grpcserver

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"route256.ozon.ru/project/loms/internal/pkg/grpc_server/middleware"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

type API interface {
	RegisterGrpcServer(server *grpc.Server)
}

type Server interface {
	Start() error
	Stop() error
	RegisterAPI(APIs []API)
}

type server struct {
	ctx        context.Context
	port       string
	grpcServer *grpc.Server
}

func NewServer(ctx context.Context, port string) Server {

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.Panic,
			middleware.Logger,
			middleware.Metrics,
			middleware.Validate,
		),
	)

	reflection.Register(grpcServer)

	return &server{
		ctx:        ctx,
		port:       port,
		grpcServer: grpcServer,
	}
}

func (s *server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("failed to create listener on port %s: %w", s.port, err)
	}

	err = s.grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to start grpc server: %w", err)
	}

	return nil
}

func (s *server) Stop() error {
	ctx := context.Background()
	s.grpcServer.GracefulStop()
	logger.Info(ctx, "grpc server is stopped successfully")
	return nil
}

func (s *server) RegisterAPI(APIs []API) {
	for _, api := range APIs {
		api.RegisterGrpcServer(s.grpcServer)
	}
}
