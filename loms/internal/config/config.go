package config

const (
	AppName                  = "loms-service"
	LomsServiceHttpPort      = "8080"
	LomsServiceGrpcPort      = "8082"
	PostgresMaster1Url       = "postgres://postgres:postgres@loms-service-master-1:5432/loms-service-1?sslmode=disable"
	PostgresSync1Url         = "postgres://postgres:postgres@loms-service-sync-1:5432/loms-service-1?sslmode=disable"
	PostgresMaster2Url       = "postgres://postgres:postgres@loms-service-master-2:5432/loms-service-2?sslmode=disable"
	PostgresSync2Url         = "postgres://postgres:postgres@loms-service-sync-2:5432/loms-service-2?sslmode=disable"
	PostgresTestUrl          = "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"
	KafkaAddr                = "kafka0:29092"
	KafkaOrderEventsTopic    = "loms.order-events"
	KafkaOutboxSenderTimeout = 5
	JaegerUrl                = "http://jaeger:14268/api/traces"
)
