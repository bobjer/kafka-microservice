# build and unite tests
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /kafka-microservice
RUN go test ./... -v

# Finalize the image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /kafka-microservice /kafka-microservice
COPY config.yaml /app/config.yaml


CMD ["/kafka-microservice"]
