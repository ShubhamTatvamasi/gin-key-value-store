# gin-key-value-store

[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/shubhamtatvamasi/gin-key-value-store?sort=semver)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Image Size (tag)](https://img.shields.io/docker/image-size/shubhamtatvamasi/gin-key-value-store/latest)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Pulls](https://img.shields.io/docker/pulls/shubhamtatvamasi/gin-key-value-store)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Go Reference](https://pkg.go.dev/badge/github.com/ShubhamTatvamasi/gin-key-value-store.svg)](https://pkg.go.dev/github.com/ShubhamTatvamasi/gin-key-value-store)
[![Go Report Card](https://goreportcard.com/badge/github.com/ShubhamTatvamasi/gin-key-value-store)](https://goreportcard.com/report/github.com/ShubhamTatvamasi/gin-key-value-store)

### Test Live Platform: https://keyvaluestore.k8s.shubhamtatvamasi.com

- [Test using Docker](#test-using-docker)
- [Test on Kubernetes](#test-on-kubernetes)
- [Test via curl](#test-via-curl)

### Test using Docker

Build docker image locally:
```bash
docker build -t shubhamtatvamasi/gin-key-value-store .
```

Run using docker:
```bash
docker run --rm -it -p 80:80 shubhamtatvamasi/gin-key-value-store
```

Run unit tests:
```bash
docker run --rm -it shubhamtatvamasi/gin-key-value-store go test
```

---

### Test on Kubernetes


create a deployment and service:
```bash
kubectl create deployment gin-key-value-store --image=shubhamtatvamasi/gin-key-value-store
kubectl expose deployment gin-key-value-store --port=80 --name=gin-key-value-store
```

Ingress value for gin-key-value-store
```bash
kubectl apply -f - << EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: gin-key-value-store
  annotations:
    cert-manager.io/cluster-issuer: letsencrypt
spec:
  tls:
  - hosts:
      - keyvaluestore.k8s.shubhamtatvamasi.com
    secretName: gin-key-value-store-tls
  rules:
  - host: keyvaluestore.k8s.shubhamtatvamasi.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: gin-key-value-store
            port:
              number: 80
EOF
```

```bash
kubectl delete deploy/gin-key-value-store svc/gin-key-value-store ing/gin-key-value-store
```

#### NodePort

Deploy on Kubernetes:
```bash
kubectl run gin-key-value-store --port=80 --expose \
  --image=shubhamtatvamasi/gin-key-value-store:0.3.2

kubectl patch svc gin-key-value-store \
  --patch='{"spec": {"type": "NodePort"}}'

kubectl patch svc gin-key-value-store \
  --patch='{"spec": {"ports": [{"nodePort": 30000, "port": 80}]}}'
```

Update the container image on pod:
```bash
kubectl set image po gin-key-value-store \
  gin-key-value-store=shubhamtatvamasi/gin-key-value-store:0.3.2
```

Delete deployment:
```bash
kubectl delete po/gin-key-value-store svc/gin-key-value-store
```

---

### Test via curl

Set key value pair:
```bash
curl -X POST -d "key=name&value=shubham" localhost/post

# Test on live server
curl -X POST -d "key=name&value=shubham" http://k8s.shubhamtatvamasi.com:30000/post
```

Get value from key store:
```bash
curl "localhost/get?key=name"

# Test on live server
curl "http://k8s.shubhamtatvamasi.com:30000/get?key=name"
```

Subscribe for a key:
```bash
curl -X POST -d "key=name" localhost/subscribe

# Test on live server
curl -X POST -d "key=name" http://k8s.shubhamtatvamasi.com:30000/subscribe
```

