# gin-key-value-store

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/shubhamtatvamasi/gin-key-value-store)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/shubhamtatvamasi/gin-key-value-store?sort=semver)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Image Size (tag)](https://img.shields.io/docker/image-size/shubhamtatvamasi/gin-key-value-store/latest)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Pulls](https://img.shields.io/docker/pulls/shubhamtatvamasi/gin-key-value-store)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![MicroBadger Layers (tag)](https://img.shields.io/microbadger/layers/shubhamtatvamasi/gin-key-value-store/latest)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/shubhamtatvamasi/gin-key-value-store)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)

Build docker image:
```bash
docker build -t key-value-store .
```

Run using docker:
```bash
docker run -d -p 80:80 shubhamtatvamasi/gin-key-value-store
```

Deploy on Kubernetes:
```bash
kubectl run gin-key-value-store --image=shubhamtatvamasi/gin-key-value-store --port=80 --expose

kubectl patch svc gin-key-value-store \
  --patch='{"spec": {"type": "NodePort"}}'

kubectl patch svc gin-key-value-store \
  --patch='{"spec": {"ports": [{"nodePort": 30000, "port": 80}]}}'
```
---

### Test Set and Get methods via curl command

Set the key
```bash
curl -X POST -d "key=name&value=shubham" localhost

# Test on live server
curl -X POST -d "key=name&value=shubham" http://k8s.shubhamtatvamasi.com:30000
```

Get the value from key store
```bash
curl "localhost/?key=name"

# Test on live server
curl "http://k8s.shubhamtatvamasi.com:30000/?key=name"
```
