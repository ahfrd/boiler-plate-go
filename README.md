# Guide

## Build image docker 
```
docker build -t ahfrd/example-boiler-plate:v1 .
```

## Run example-boiler-plate image on docker
```
docker run -d -p 9018:9018 -v config:/app/config --name example-boiler-plate-v1 ahfrd/example-boiler-plate:v1
```# AQ
