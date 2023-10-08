ARG SPARK_BASE_IMAGE
ARG SPARK_APP_EXAMPLE_IMAGE

FROM $SPARK_APP_EXAMPLE_IMAGE AS BUILD

FROM $SPARK_BASE_IMAGE AS RUNTIME

USER root

ENV SPARK_HOME="/opt/spark" \
    HADOOP_CONF_DIR="${SPARK_HOME}/conf" \
    PATH="${SPARK_HOME}/bin:${SPARK_HOME}/sbin/:/opt/:${PATH}"

RUN apt-get update && apt-get install -y curl \
    && curl https://dl.min.io/client/mc/release/linux-amd64/mc -o /opt/mc \
    && chmod +x /opt/mc

COPY --from=BUILD /opt/spark-app-example/spark-app-example.jar ${SPARK_HOME}/work-dir/
COPY entrypoint.sh /opt/spark/sbin

RUN chmod +x /opt/spark/sbin/*.sh \
    && chmod +x /opt/spark/bin/*.sh

ENTRYPOINT ["entrypoint.sh"]