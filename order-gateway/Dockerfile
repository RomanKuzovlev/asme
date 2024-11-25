FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o order-gateway ./cmd/.

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/order-gateway .

EXPOSE 8080

ENTRYPOINT ["./order-gateway"]
