FROM python:3.9.18-slim

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

COPY entrypoint.sh /opt/airflow/entrypoint.sh
RUN chmod +x /opt/airflow/entrypoint.sh

ENTRYPOINT ["/opt/airflow/bootstrap.sh"]