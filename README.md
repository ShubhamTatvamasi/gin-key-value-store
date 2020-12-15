# gin-key-value-store

[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/shubhamtatvamasi/gin-key-value-store)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/shubhamtatvamasi/gin-key-value-store?sort=semver)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Image Size (tag)](https://img.shields.io/docker/image-size/shubhamtatvamasi/gin-key-value-store/latest)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Pulls](https://img.shields.io/docker/pulls/shubhamtatvamasi/gin-key-value-store)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/shubhamtatvamasi/gin-key-value-store)](https://hub.docker.com/r/shubhamtatvamasi/gin-key-value-store)

### Test using Docker

Build docker image locally:
```bash
docker build -t shubhamtatvamasi/gin-key-value-store .
```

Run using docker:
```bash
docker run --rm -it -p 80:80 shubhamtatvamasi/gin-key-value-store
```

### Test on Kubernetes

Deploy on Kubernetes:
```bash
kubectl run gin-key-value-store --port=80 --expose \
  --image=shubhamtatvamasi/gin-key-value-store:0.1.0

kubectl patch svc gin-key-value-store \
  --patch='{"spec": {"type": "NodePort"}}'

kubectl patch svc gin-key-value-store \
  --patch='{"spec": {"ports": [{"nodePort": 30000, "port": 80}]}}'
```

Update the container image on pod:
```bash
kubectl set image po gin-key-value-store gin-key-value-store=shubhamtatvamasi/gin-key-value-store:0.2.0
```

Delete deployment
```bash
kubectl delete po/gin-key-value-store svc/gin-key-value-store
```

---

### Test Set and Get methods via curl command

Set key value pair
```bash
curl -X POST -d "key=name&value=shubham" localhost/post

# Test on live server
curl -X POST -d "key=name&value=shubham" http://k8s.shubhamtatvamasi.com:30000/post
```

Get value from key store
```bash
curl "localhost/get?key=name"

# Test on live server
curl "http://k8s.shubhamtatvamasi.com:30000/get?key=name"
```

Subscribe for a key
```bash
curl -X POST -d "key=name" localhost/subscribe

# Test on live server
curl -X POST -d "key=name" http://k8s.shubhamtatvamasi.com:30000/subscribe
```
