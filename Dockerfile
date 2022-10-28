FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod tidy
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8080

ENTRYPOINT go run main.go