#!/bin/bash


export MINIO_DOMAIN="s3"
export MINIO_ROOT_USER="minioadmin"
export MINIO_ROOT_PASSWORD="minioadmin"
export MINIO_ACCESS_KEY="minioadmin"
export MINIO_SECRET_KEY="minioadmin"

export S3_ACCESS_KEY="minioadmin"
export S3_SECRET_KEY="minioadmin"

if [[ "$1" == "serve" ]] 
then
    echo "Running Minio Server"
    minio server /data --address ":$PORT" --console-address ":$CONSOLE_PORT"

elif [[ "$1" == "init" ]] 
then
    echo "Waiting for server to startup ..."
    sleep 5
    echo "Creating s3 bucket"
    mc config host add default_server http://$ENDPOINT_URL:$PORT $S3_ACCESS_KEY $S3_SECRET_KEY --api s3v4;
    mc mb default_server/$S3_BUCKET_NAME;
    echo "mc is configured !"
fi
