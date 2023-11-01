# grpc-go-sample

## Prepare
```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt --fix-broken install
sudo apt install golang

VER=24.1
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${VER}/protoc-${VER}-linux-x86_64.zip
sudo unzip protoc-${VER}-linux-x86_64.zip -d /usr/local/protobuf

RUNTIME_VER=v1.31.0
curl -OL https://github.com/protocolbuffers/protobuf-go/releases/download/v1.31.0/protoc-gen-go.${RUNTIME_VER}.linux.amd64.tar.gz
sudo tar xzvf protoc-gen-go.${RUNTIME_VER}.linux.amd64.tar.gz -C /usr/local/protobuf/bin

export PATH=$PATH:/usr/local/protobuf/bin
echo 'export PATH=$PATH:/usr/local/protobuf/bin' >> ~/bashrc

export PATH=$PATH:$HOME/go/bin
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/bashrc
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Build
```
cd example
export PATH=$PATH:/usr/local/go/bin
make 
docker build -t <your repo> . 
docker push <your repo>
```
## Run
Edit IMAGE in the script.
```
vi run.sh
```
Run the script.
```
./run.sh
```

## Load test

```js
cat << EOF > ./k6.js
import http from 'k6/http';

export let options = {
    discardResponseBodies: true,
    scenarios: {
      test_rate1: {
        executor: 'constant-arrival-rate',
        duration: '30s',
        rate: 2000,
        timeUnit: '1s',
        preAllocatedVUs: 1000,
        maxVUs: 1000,
        },
      },
    };

export default function () {
  http.get('http://grpc-client:8080/');
}
EOF
```
```
kubectl create configmap k6-test --from-file ./k6.js
```
```yaml
cat << EOF > ./k6-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: k6
  name: k6-sample
spec:
  containers:
  - command:
    - k6
    - run
    - /test/k6.js
    image: ghcr.io/grafana/k6
    name: k6
    ports:
    volumeMounts:
    - mountPath: /test
      name: k6-test-volume
  volumes:
  - configMap:
      name: k6-test
    name: k6-test-volume
EOF
```

```
kubectl apply -f ./k6-pod.yaml
```
