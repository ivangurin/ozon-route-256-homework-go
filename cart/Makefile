run:
	go run .

test:
	go test ./...

build:
	go build -o ./bin/cart .

docker-run:
	docker build -t cart .
	docker run -d -p 8080:8080 --rm --name cart cart

docker-stop:
	docker stop cart
	docker rmi cart

