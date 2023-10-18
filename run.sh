#!/bin/bash
set -x

IMAGE=imuratashared/grpc-server

kubectl run --image $IMAGE grpc-client 
kubectl create deploy --image $IMAGE grpc-server -r 1 --port 9000
kubectl wait --timeout=30s --for=condition=Ready pod grpc-client
kubectl wait --timeout=30s --for=condition=available deploy grpc-server
kubectl expose pod grpc-client --port 8080
kubectl expose deploy grpc-server --port 9000
kubectl exec -it grpc-client -- bash -c './client --server grpc-server:9000'
#for i in {1..10000}; do
#kubectl exec -it grpc-client -- bash -c 'for i in {1..10000}; do ./client --server grpc-server:9000; done'
#done





