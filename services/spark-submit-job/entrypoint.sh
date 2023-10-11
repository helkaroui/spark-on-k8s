#!/bin/bash

echo "Wait for S3 server to be available"
until [ "$(curl -s -w '%{http_code}' -o /dev/null "http://s3:9000/minio/health/live")" -eq 200 ]
do
    echo "Waiting for server to startup ..."
    sleep 3
done


export S3_ACCESS_KEY="minioadmin"
export S3_SECRET_KEY="minioadmin"
export S3_BUCKET_NAME=dev-default-bucket


echo "Configure Minio client"
mc config host add default_server http://s3:9000 $S3_ACCESS_KEY $S3_SECRET_KEY --api s3v4;

echo "Creating s3 bucket"
mc mb default_server/$S3_BUCKET_NAME;

echo "Upload Jar to s3"
mc cp ./spark-app-example.jar default_server/$S3_BUCKET_NAME/





K8S_API_TOKEN="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)"
K8S_API_CERT="/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
SPARK_DOCKER_IMAGE="spark-base-image:latest"

spark-submit \
    --master k8s://https://kubernetes.default.svc:443 \
    --deploy-mode cluster \
    --name spark-pi \
    --class dev.sharek.examples.SparkPi \
    --conf spark.executor.instances=2 \
    --conf spark.driver.extraJavaOptions="--add-exports java.base/sun.nio.ch=ALL-UNNAMED" \
    --supervise \
    --conf spark.kubernetes.container.image=$SPARK_DOCKER_IMAGE \
    --conf spark.kubernetes.authenticate.driver.serviceAccountName=$SERVICE_ACCOUNT_NAME \
    --conf spark.kubernetes.authenticate.executor.serviceAccountName=$SERVICE_ACCOUNT_NAME \
    --conf spark.kubernetes.authenticate.submission.oauthToken=$K8S_API_TOKEN \
    --conf spark.kubernetes.authenticate.submission.caCertFile=$K8S_API_CERT \
    --conf spark.hadoop.fs.s3a.endpoint=http://s3:9000 \
    --conf spark.hadoop.fs.s3a.path.style.access=true \
    --conf spark.hadoop.fs.s3a.access.key=$S3_ACCESS_KEY \
    --conf spark.hadoop.fs.s3a.secret.key=$S3_SECRET_KEY \
    --conf spark.ui.reverseProxy=true \
    s3a://$S3_BUCKET_NAME/spark-app-example.jar 10 600
