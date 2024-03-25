build-all:
	cd cart && GOOS=linux GOARCH=amd64 make build
	cd loms && GOOS=linux GOARCH=amd64 make build

run-all: build-all
	docker-compose up --force-recreate --build

run:
	docker-compose up -d --force-recreate
	cd loms && make migrate

stop:
	docker-compose down
	docker rmi cart loms
