FROM golang:1.24.2-alpine AS builder

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN swag init

RUN go build -o go-server


CMD ["/app/go-server"]
