package config

const (
	AppName               = "cart-service"
	CartServiceHttpPort   = "8080"
	ProductServiceHost    = "route256.pavl.uk:8080"
	ProductServiceToken   = "testtoken"
	ProductServiceRetries = 3
	LomsServiceGrpcHost   = "loms:8082"
	JaegerUrl             = "http://jaeger:14268/api/traces"
	RedisUrl              = "redis://redis:6379/0?protocol=3"
)
