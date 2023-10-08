#!/bin/bash

export MINIO_DOMAIN="s3"
export MINIO_ROOT_USER="minioadmin"
export MINIO_ROOT_PASSWORD="minioadmin"
export MINIO_ACCESS_KEY="minioadmin"
export MINIO_SECRET_KEY="minioadmin"

export S3_ACCESS_KEY="minioadmin"
export S3_SECRET_KEY="minioadmin"


echo "Running Minio Server"
minio server /data --address ":$PORT" --console-address ":$CONSOLE_PORT"