#!/bin/bash

IMAGE_NAME=fanyaoyao12138/wechat_monitor:latest
DOCKER_CONTAINER=wechat_monitor

echo "stop rm docker container"
docker stop $DOCKER_CONTAINER
docker rm $DOCKER_CONTAINER

echo "run docker container"
docker run -p 8088:8088 -d --restart=always --name $DOCKER_CONTAINER $IMAGE_NAME

echo "docker log"
docker logs -f --tail=100 $DOCKER_CONTAINER