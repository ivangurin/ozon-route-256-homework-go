# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

run:
	go run ./cmd/cart_service

test:
	go test -v -race -count 1 ./...

bench:
	go test -v -bench=. -benchmem ./...

build:
	go build -o ./bin/cart_service ./cmd/cart_service

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

generate: genproto genmock

genmock:
	$(info $(shell printf "\033[34;1m▶\033[0m") go generate-mocks...)
	@for f in $(shell find internal -name 'genmock.go'| sort -u); do \
		go generate $$f; \
	done

PHONY: genproto
genproto: .proto-generate
	go mod tidy

PHONY: .proto-generate
.proto-generate: .bin-deps .vendor-proto .order-api-generate .stock-api-generate

# Вендоринг внешних proto файлов
.vendor-proto: .vendor-rm  vendor-proto/google/protobuf vendor-proto/validate vendor-proto/google/api vendor-proto/protoc-gen-openapiv2/options
	go mod tidy

.PHONY: .vendor-rm
.vendor-rm:
	rm -rf vendor-proto

.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)
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
	rm -rf internal/pb/${ORDER_API_PROTO_PATH}
	mkdir -p internal/pb/${ORDER_API_PROTO_PATH}
	protoc \
	-I ../loms/${ORDER_API_PROTO_PATH} \
	-I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go_out internal/pb/${ORDER_API_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--go-grpc_out internal/pb/${ORDER_API_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
	--validate_out="lang=go,paths=source_relative:internal/pb/${ORDER_API_PROTO_PATH}" \
	../loms/${ORDER_API_PROTO_PATH}/*.proto

STOCK_API_PROTO_PATH:=api/stock/v1
PHONY: .stock-api-generate
.stock-api-generate: 
	rm -rf internal/pb/${STOCK_API_PROTO_PATH}
	mkdir -p internal/pb/${STOCK_API_PROTO_PATH}
	protoc \
	-I ../loms/${STOCK_API_PROTO_PATH} \
	-I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go \
	--go_out internal/pb/${STOCK_API_PROTO_PATH} \
	--go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc \
	--go-grpc_out internal/pb/${STOCK_API_PROTO_PATH} \
	--go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate \
	--validate_out="lang=go,paths=source_relative:internal/pb/${STOCK_API_PROTO_PATH}" \
	../loms/${STOCK_API_PROTO_PATH}/*.proto

docker-build:
	@docker build -t cart .
	@docker image prune -f --filter label=stage=builder

docker-run: docker-build
	@docker run -d -p 8082:8080 --rm --name cart cart

docker-stop:
	@docker stop cart || true
	@docker rmi cart || true

prepare-integration-test:

integration-test-run: prepare-integration-test
	go test -v -race -tags integration ./...

integration-test: docker-run integration-test-run docker-stop
	