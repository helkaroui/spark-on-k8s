FROM golang:1.17.3-alpine

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

WORKDIR /opt/source

COPY . /opt/source/

RUN go mod download \
    && go build -a -o /opt/spark-ui-reverse-proxy main.go
