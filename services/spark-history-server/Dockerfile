ARG SPARK_BASE_IMAGE

FROM ${SPARK_BASE_IMAGE} as RUNTIME

COPY spark-defaults.conf ${SPARK_HOME}/conf
COPY --chown=185:185 entrypoint.sh ${SPARK_HOME}/sbin

USER root

RUN chmod +x ${SPARK_HOME}/sbin/entrypoint.sh \
    && mkdir -p /opt/spark/logs/ \
    && chown -R 185:185 /opt/spark/logs/

USER 185

ENV SPARK_HOME="/opt/spark" \
    HADOOP_CONF_DIR="${SPARK_HOME}/conf" \
    PATH="${SPARK_HOME}/bin:${SPARK_HOME}/sbin/:/opt/:${PATH}"

ENTRYPOINT [ "entrypoint.sh" ]
