FROM golang:alpine as builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o ./bin/cart .

FROM alpine

COPY --from=builder /app/bin/cart /cart

CMD [ "/cart" ]
