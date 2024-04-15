package middleware

import (
	"net/http"
	"time"

	"route256.ozon.ru/project/cart/internal/pkg/metrics"
)

func WithMetrics(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			metrics.UpdateRequestsTotal(r.Method, r.URL.Path)
			defer metrics.UpdateResponseTime(time.Now().UTC())
			next.ServeHTTP(w, r)
			lw := NewLoggingResponseWriter(w)
			metrics.UpdateResponseCode(r.Method, r.URL.Path, http.StatusText(lw.statusCode))
		},
	)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
