# Guide

## Build image docker 
```
docker build -t ahfrd/asia-quest:v1 .
```

## Run asia-quest image on docker
```
docker run -d -p 9018:9018 -v config:/app/config --name asia-quest-v1 ahfrd/asia-quest:v1
```# AQ
