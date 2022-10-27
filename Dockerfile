# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /go-rest-api

EXPOSE 8080

CMD [ "/go-rest-api" ]

RUN go run main.go

