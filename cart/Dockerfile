FROM golang as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/cart_service ./cmd/cart_service

FROM ubuntu

COPY --from=builder /app/bin/cart_service /cart_service

CMD [ "/cart_service" ]
