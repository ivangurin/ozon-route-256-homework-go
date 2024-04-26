package middleware

import (
	"fmt"
	"net/http"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
	"route256.ozon.ru/project/cart/internal/pkg/tracer"
)

func WithTracer(next http.Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.StartSpanFromContext(r.Context(), fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		defer span.End()
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func WithLogger(next http.Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger.Infof(r.Context(), "start to handle request method: %s, url: %s", r.Method, r.URL.Path)
		defer logger.Infof(r.Context(), "finished handle request method: %s, url: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
