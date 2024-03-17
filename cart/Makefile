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

genmock:
	$(info $(shell printf "\033[34;1m▶\033[0m") go generate-mocks...)
	@for f in $(shell find internal -name 'genmock.go'| sort -u); do \
		go generate $$f; \
	done

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
	