package middleware

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"route256.ozon.ru/project/loms/internal/pkg/logger"
)

func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	raw, _ := protojson.Marshal((req).(proto.Message)) // для превращения protobuf структур в json используем google.golang.org/protobuf/encoding/protojson пакет а не encoding/json
	logger.Infof(ctx, "request: method: %v, req: %v", info.FullMethod, string(raw))

	if resp, err = handler(ctx, req); err != nil {
		logger.Infof(ctx, "response: method: %v, err: %v", info.FullMethod, err)
		return
	}

	rawResp, _ := protojson.Marshal((resp).(proto.Message))
	logger.Infof(ctx, "response: method: %v, resp: %v", info.FullMethod, string(rawResp))

	return
}

func WithHTTPLoggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method)
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
