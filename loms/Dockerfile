FROM golang as builder
LABEL stage=builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/loms_service ./cmd/loms_service

FROM ubuntu

COPY --from=builder /app/bin/loms_service /loms_service

CMD [ "/loms_service" ]
