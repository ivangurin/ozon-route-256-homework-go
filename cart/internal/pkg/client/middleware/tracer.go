package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func Tracer(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return nil
}
