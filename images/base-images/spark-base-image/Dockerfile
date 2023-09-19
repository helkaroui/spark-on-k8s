FROM maven:3.9.4-eclipse-temurin-8-focal AS BUILD

ENV SRC_DIR=/opt/source/
ENV DIST_DIR=/opt/spark/

ARG SPARK_VERSION

ENV SPARK_VERSION=${SPARK_VERSION}

RUN apt-get update \
    && apt-get install -y unzip

RUN mkdir -p ${SRC_DIR} \
    && mkdir -p ${DIST_DIR} \
    && cd ${SRC_DIR} \
    && curl -skL -XGET https://github.com/apache/spark/archive/refs/tags/v${SPARK_VERSION}.zip -o ${SRC_DIR}/spark.zip \
    && unzip -qq spark.zip \
    && cd spark-${SPARK_VERSION} \
    && mvn install -pl core,streaming,assembly -Pkubernetes -Phive -Phive-thriftserver -T 4 -DskipTests --no-transfer-progress \
    && cp -r sbin bin conf assembly/target/scala-2.12/jars ${DIST_DIR} \
    && mvn clean \
    && rm -rf ${SRC_DIR}

FROM eclipse-temurin:8-focal as RUNTIME

ARG spark_uid=185

RUN set -ex && \
    apt-get update && \
    ln -s /lib /lib64 && \
    apt install -y bash tini && \
    mkdir -p /opt/spark && \
    mkdir -p /opt/spark/examples && \
    mkdir -p /opt/spark/work-dir && \
    touch /opt/spark/RELEASE && \
    rm /bin/sh && \
    ln -sv /bin/bash /bin/sh && \
    echo "auth required pam_wheel.so use_uid" >> /etc/pam.d/su && \
    chgrp root /etc/passwd && chmod ug+rw /etc/passwd && \
    rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/*

ENV SPARK_HOME="/opt/spark" \
    HADOOP_CONF_DIR="${SPARK_HOME}/conf" \
    PATH="${SPARK_HOME}/bin:${SPARK_HOME}/sbin/:/opt/:${PATH}"

COPY --from=BUILD --chown=185:185 /opt/spark ${SPARK_HOME}
COPY --chown=185:185 *.sh ${SPARK_HOME}/sbin/

RUN chmod g+w ${SPARK_HOME}/work-dir \
      && chmod a+x ${SPARK_HOME}/sbin/*.sh \
      && chmod a+x ${SPARK_HOME}/bin/*.sh

WORKDIR ${SPARK_HOME}/work-dir

ENTRYPOINT [ "entrypoint.sh" ]

# Specify the User that the actual main process will run as
USER ${spark_uid}