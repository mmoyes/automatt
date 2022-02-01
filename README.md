# AutoMatt
A rest api to help automate some of the things Matt does


## run in docker for dev
```
go build && \
docker build -t automatt:latest . && \
docker run \
  -p 8800:8800 \
  --mount type=bind,source="$(pwd)"/doConfig.yml,target=/home/user/.kube/config \
  automatt:latest

```

## deploy

```
go build && \
docker build -t bam.brud.local:6000/brudtech/automatt:latest . && \
docker push bam.brud.local:6000/brudtech/automatt:latest && \
kubectl patch deployment automatt-main-dpy -p \
  "{\"spec\":{\"template\":{\"metadata\":{\"annotations\":{\"date\":\"`date +'%s'`\"}}}}}"
```

deploy.yml and secret