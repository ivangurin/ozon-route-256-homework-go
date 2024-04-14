package middleware

import (
	"net/http"

	"route256.ozon.ru/project/cart/internal/pkg/logger"
)

func WithLogger(next http.Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		logger.Infof(r.Context(), "start to handle request method: %s, url: %s", r.Method, r.URL.Path)
		defer logger.Infof(r.Context(), "finished handle request method: %s, url: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
