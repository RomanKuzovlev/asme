FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o matching-engine ./cmd/.

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/matching-engine .
COPY .env .

EXPOSE 50051

ENTRYPOINT ["./matching-engine"]
