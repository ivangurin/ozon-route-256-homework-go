run:
	go run ./cmd/cart_service

test:
	go test -v -race ./...

build:
	go build -o ./bin/cart_service ./cmd/cart_service

docker-build:
	docker build -t cart .
	docker image prune -f --filter label=stage=builder

docker-run: docker-build
	docker run -d -p 8082:8080 --rm --name cart cart

docker-stop:
	docker stop cart || true
	docker rmi cart || true

generate-mocks:
	$(info $(shell printf "\033[34;1m▶\033[0m") go generate-mocks...)
	@for f in $(shell find internal -name 'genmock.go'| sort -u); do \
		go generate $$f; \
	done