FROM sbtscala/scala-sbt:eclipse-temurin-jammy-8u382-b05_1.9.6_3.3.1

ENV SRC_DIR=/opt/source
ENV APP_DIR=/opt/spark-app-example


RUN mkdir -p $SRC_DIR \
    && mkdir -p $APP_DIR

COPY . $SRC_DIR

RUN cd $SRC_DIR \
    && sbt 'set test in assembly := {}' clean assembly \
    && cp target/scala-3.3.1/spark-app-example.jar $APP_DIR/

WORKDIR $APP_DIR
