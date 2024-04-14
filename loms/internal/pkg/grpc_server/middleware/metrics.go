package middleware

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256.ozon.ru/project/loms/internal/pkg/metrics"
)

func Metrics(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	metrics.UpdateRequestsCounter(info.FullMethod)

	defer metrics.UpdateResponseTime(time.Now())

	resp, err = handler(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		metrics.UpdateResponseCode(info.FullMethod, st.Code().String())
	} else {
		metrics.UpdateResponseCode(info.FullMethod, codes.OK.String())
	}

	return
}
