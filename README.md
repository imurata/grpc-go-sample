# grpc-go-sample

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
