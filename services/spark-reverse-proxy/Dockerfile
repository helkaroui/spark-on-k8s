FROM golang:1.20-alpine AS BUILD

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on

WORKDIR /opt/source

COPY go.mod /opt/source/go.mod
COPY go.sum /opt/source/go.sum

RUN go mod download
COPY . /opt/source/
RUN go build -a -o /opt/spark-ui-reverse-proxy main.go

FROM alpine:3.14 as RUNTIME
USER root

COPY --from=BUILD /opt/spark-ui-reverse-proxy /usr/bin/
COPY --from=BUILD /opt/source/templates /templates

ENV KUBERNETES_SERVICE_HOST="kubernetes.default.svc" \
    KUBERNETES_SERVICE_PORT="443" \
    GIN_MODE=release

ENTRYPOINT ["/usr/bin/spark-ui-reverse-proxy"]
