ARG REVERSE_PROXY_IMAGE

FROM ${REVERSE_PROXY_IMAGE} as BUILD

FROM alpine:3.14 as RUNTIME
USER root

COPY --from=BUILD /opt/spark-ui-reverse-proxy /usr/bin/

ENTRYPOINT ["/usr/bin/spark-ui-reverse-proxy"]
