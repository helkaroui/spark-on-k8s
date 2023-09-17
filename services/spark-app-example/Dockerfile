ARG SPARK_BASE_IMAGE
ARG SPARK_APP_EXAMPLE_IMAGE

FROM $SPARK_APP_EXAMPLE_IMAGE AS BUILD

FROM $SPARK_BASE_IMAGE AS RUNTIME

COPY --from=BUILD /opt/spark-app-example/spark-app-example.jar ${SPARK_HOME}/work-dir/
COPY entrypoint.sh /opt/spark/sbin

RUN chmod +x /opt/spark/sbin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]