FROM golang:1.24.2-alpine AS builder

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.4

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go list &&\
    go test &&\
#    swag init &&\
    GOOS=linux go build -o go-server
#RUN swag init

#RUN go build -o go-server


CMD ["/app/go-server"]
