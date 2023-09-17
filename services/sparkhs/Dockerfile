ARG SPARK_BASE_IMAGE

FROM ${SPARK_BASE_IMAGE} as RUNTIME

COPY spark-defaults.conf ${SPARK_HOME}/conf
COPY entrypoint.sh ${SPARK_HOME}/sbin

RUN chmod +x ${SPARK_HOME}/sbin/entrypoint.sh

ENTRYPOINT [ "entrypont.sh" ]
