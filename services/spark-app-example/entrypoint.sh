#!/bin/bash

K8S_ENDPOINT="https://<k8s-apiserver-host>:<k8s-apiserver-port>"
DRIVER_NAME=""

spark-submit \
    --master k8s://$K8S_ENDPOINT \
    --deploy-mode cluster \
    --name spark-pi \
    --class dev.sharek.examples.SparkPi \
    --conf spark.executor.instances=2 \
    --conf spark.kubernetes.container.image=<spark-image> \
    local:///opt/spark/work-dir/spark-app-example.jar