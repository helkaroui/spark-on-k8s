FROM python:3.9.18-slim AS RUNTIME

# install deps
RUN apt-get update -y && apt-get install -y \
    libczmq-dev \
    libssl-dev \
    inetutils-telnet \
    bind9utils \
    gcc \
    && apt-get clean

RUN pip install --upgrade pip

RUN pip install apache-airflow==2.7.1
RUN pip install 'apache-airflow[kubernetes]'

WORKDIR /opt/airflow

COPY entrypoint.sh /opt/airflow/
COPY airflow.cfg /opt/airflow/
COPY airflow.db /opt/airflow/
COPY dags /opt/airflow/dags/

RUN chmod +x /opt/airflow/entrypoint.sh \
    && mkdir -p /opt/airflow/logs

ENV AIRFLOW_HOME="/opt/airflow/" \
    DAG_FOLDER="/opt/airflow/dags"

ENTRYPOINT ["/opt/airflow/entrypoint.sh"]