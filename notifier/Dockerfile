FROM golang as builder
LABEL stage=builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/notifier_service ./cmd/notifier_service

FROM ubuntu

COPY --from=builder /app/bin/notifier_service /notifier_service

CMD [ "/notifier_service" ]
