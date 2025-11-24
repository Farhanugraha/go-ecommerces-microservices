FROM golang:1.25.4 AS builder

WORKDIR /app

COPY . .

RUN go build -o auth-service ./cmd/main.go

FROM debian:stable-slim

WORKDIR /app

COPY --from=builder /app/auth-service .

EXPOSE 8080

CMD ["./auth-service"]
