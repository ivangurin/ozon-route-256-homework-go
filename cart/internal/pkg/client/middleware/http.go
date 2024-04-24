package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"route256.ozon.ru/project/cart/internal/pkg/metrics"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func NewMetricInterceptor(serviceName string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		metrics.UpdateExternalRequestsTotal(
			serviceName,
			method,
		)
		defer metrics.UpdateExternalResponseTime(time.Now().UTC())

		err := invoker(ctx, method, req, reply, cc, opts...)

		if err != nil {
			st, _ := status.FromError(err)
			metrics.UpdateExternalResponseCode(
				serviceName,
				method,
				st.Code().String(),
			)

			return err
		}

		metrics.UpdateExternalResponseCode(
			serviceName,
			method,
			codes.OK.String(),
		)

		return nil
	}
}

type MetricRoundTripper struct {
	Proxied     http.RoundTripper
	serviceName string
}

func NewHttpMiddleware(serviceName string) *MetricRoundTripper {
	return &MetricRoundTripper{
		Proxied:     http.DefaultTransport,
		serviceName: serviceName,
	}
}

func (mrt *MetricRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	ctx, span := tracer.StartSpanFromContext(r.Context(), fmt.Sprintf("%s%s", mrt.serviceName, r.URL.Path))
	defer span.End()
	r = r.WithContext(ctx)

	metrics.UpdateExternalRequestsTotal(
		mrt.serviceName,
		r.URL.Path,
	)
	defer metrics.UpdateExternalResponseTime(time.Now().UTC())

	resp, err := mrt.Proxied.RoundTrip(r)

	metrics.UpdateExternalResponseCode(
		mrt.serviceName,
		r.URL.Path,
		http.StatusText(resp.StatusCode),
	)

	return resp, err
}
