#!/bin/bash
set -x

IMAGE_CLIENT=imuratashared/grpc-client
IMAGE_SERVER=imuratashared/grpc-server

kubectl create deploy --image $IMAGE grpc-server -r 1 --port 9000
kubectl wait --timeout=30s --for=condition=available deploy grpc-server
kubectl expose deploy grpc-server --port 9000
kubectl run --image $IMAGE_CLIENT grpc-client -- --server grpc-server:9000
kubectl wait --timeout=30s --for=condition=Ready pod grpc-client
kubectl expose pod grpc-client --port 8080





