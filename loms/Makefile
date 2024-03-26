# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

run:
	go run ./cmd/loms_service

test:
	go test -v -race -count 1 ./...

build:
	go build -o ./bin/loms_service ./cmd/loms_service

lint:
	golangci-lint run \
	--config=.golangci.yaml \
	./...

cover:
	@go test ./... -coverprofile cover.out.tmp > /dev/null
	@grep -vE "_mock.go" cover.out.tmp > cover.out
	@rm -f cover.out.tmp
	@go tool cover -func cover.out | fgrep total | awk '{print "Percent of coverage is: " $$3}'
	@rm -f cover.out

cover-html:
	@go test ./... -coverprofile covertmp.out > /dev/null
	@grep -vE "_mock.go" covertmp.out > cover.out
	@rm -f covertmp.out
	@go tool cover -html cover.out 
	@rm -f cover.out

generate: genproto genmock gensqlc

PHONY: .bin-minimock
.bin-minimock:
	$(info $(shell printf "\033[34;1m▶\033[0m") Installing minimock...)
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest

genmock: .bin-minimock
	$(info $(shell printf "\033[34;1m▶\033[0m") Generating mocks...)
	@for f in $(shell find internal -name 'genmock.go'| sort -u); do \
		PATH=$(LOCAL_BIN):"$(PATH)" go generate $$f; \
	done

PHONY: genproto
genproto: .proto-generate
	go mod tidy

PHONY: .proto-generate
.proto-generate: .bin-proto .vendor-proto  .order-api-generate .stock-api-generate .merge-swagger

# Вендоринг внешних proto файлов
.vendor-proto: .vendor-rm  vendor-proto/google/protobuf vendor-proto/validate vendor-proto/google/api vendor-proto/protoc-gen-openapiv2/options
	go mod tidy

.PHONY: .vendor-rm
.vendor-rm:
	rm -rf vendor-proto

.PHONY: .bin-proto
.bin-proto:
	$(info $(shell printf "\033[34;1m▶\033[0m") Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
    GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.5
	GOBIN=$(LOCAL_BIN) go install github.com/ntauth/go-swagger-merger@v0.1.0

# Устанавливаем proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

# Устанавливаем proto описания validate
vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor-proto/validate
		mv vendor-proto/tmp/validate vendor-proto/
		rm -rf vendor-proto/tmp

# Устанавливаем proto описания google/googleapis
vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis && \
 	cd vendor-proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

# Устанавливаем proto описания protoc-gen-openapiv2/options
vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/grpc-ecosystem && \
 	cd vendor-proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/grpc-ecosystem

ORDER_API_PROTO_PATH:=api/order/v1
PHONY: .order-api-generate
.order-api-generate: 
	rm -rf pkg/${ORDER_API_PROTO_PATH}
	mkdir -p pkg/${ORDER_API_PROTO_PATH}
	protoc \
	-I ${ORDER_API_PROTO_PATH} \
	-I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go_out pkg/${ORDER_API_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--go-grpc_out pkg/${ORDER_API_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
	--validate_out="lang=go,paths=source_relative:pkg/${ORDER_API_PROTO_PATH}" \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--grpc-gateway_out pkg/${ORDER_API_PROTO_PATH} \
	--grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
	--openapiv2_out pkg/${ORDER_API_PROTO_PATH} \
	--openapiv2_opt logtostderr=true \
	${ORDER_API_PROTO_PATH}/*.proto

STOCK_API_PROTO_PATH:=api/stock/v1
PHONY: .stock-api-generate
.stock-api-generate: 
	rm -rf pkg/${STOCK_API_PROTO_PATH}
	mkdir -p pkg/${STOCK_API_PROTO_PATH}
	protoc \
	-I ${STOCK_API_PROTO_PATH} \
	-I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go_out pkg/${STOCK_API_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--go-grpc_out pkg/${STOCK_API_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
	--validate_out="lang=go,paths=source_relative:pkg/${STOCK_API_PROTO_PATH}" \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	--grpc-gateway_out pkg/${STOCK_API_PROTO_PATH} \
	--grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true \
	--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 \
	--openapiv2_out pkg/${STOCK_API_PROTO_PATH} \
	--openapiv2_opt logtostderr=true \
	${STOCK_API_PROTO_PATH}/*.proto

PHONY: .merge-swagger
.merge-swagger: 
	rm -rf pkg/swagger
	mkdir -p pkg/swagger
	$(LOCAL_BIN)/go-swagger-merger \
	-o pkg/swagger/swagger.json \
	pkg/${ORDER_API_PROTO_PATH}/order.swagger.json \
	pkg/${STOCK_API_PROTO_PATH}/stock.swagger.json


GOOSE_DRIVER=postgres
GOOSE_DBSTRING="host=localhost user=postgres password=postgres dbname=loms-service sslmode=disable"
GOOSE_MIGRATION_DIR=./migrations

PHONY: .bin-goose
.bin-goose:
	$(info $(shell printf "\033[34;1m▶\033[0m") Installing goose binary...)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

migrate: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

migrate-down: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

migrate-status: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) status

PHONY: .bin-sqlc
.bin-sqlc:
	$(info $(shell printf "\033[34;1m▶\033[0m") Installing sqlc binary...)
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

gensqlc: .bin-sqlc
	$(LOCAL_BIN)/sqlc generate

GOOSE_DBSTRING_TEST="host=localhost port=54321 user=postgres password=postgres dbname=postgres sslmode=disable"

integration-test:  integration-test-up integration-test-migrate integration-test-run integration-test-down
	
integration-test-up: 
	docker-compose -f ./internal/pkg/suite/integration_test/docker-compose-test.yml up -d

integration-test-migrate: .bin-goose
	$(LOCAL_BIN)/goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING_TEST) up || true

integration-test-run:
	go test -v -tags integration ./... || true

integration-test-down:
	docker-compose -f ./internal/pkg/suite/integration_test/docker-compose-test.yml down
