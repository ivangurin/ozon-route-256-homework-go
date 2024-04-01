package config

const (
	LomsServiceHttpPort = "8080"
	LomsServiceGrpcPort = "8082"
	MasterDBUrl         = "postgres://postgres:postgres@loms-service-master:5432/loms-service?sslmode=disable"
	SyncDBUrl           = "postgres://postgres:postgres@loms-service-sync:5432/loms-service?sslmode=disable"
	TestDBUrl           = "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"
)
