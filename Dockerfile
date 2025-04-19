FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-server .
EXPOSE 8080

CMD ["/app/go-server"]
