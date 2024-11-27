#!/bin/bash

# 清除所有 none 镜像
echo "Removing all dangling <none> images..."
docker images -f "dangling=true" -q | xargs -r docker rmi

# 定义需要重新构建的服务名称
SERVICES=()

# 停止并删除旧的容器
for SERVICE in "${SERVICES[@]}"; do
  echo "Stopping and removing old container for $SERVICE..."
  docker-compose -f docker-compose.yml rm -sf $SERVICE
done

# 删除旧镜像
for SERVICE in "${SERVICES[@]}"; do
  IMAGE_NAME=$(docker-compose -f docker-compose.yml config | grep "image: " | grep "$SERVICE" | awk '{print $2}')
  if [ ! -z "$IMAGE_NAME" ]; then
    echo "Removing old image $IMAGE_NAME..."
    docker rmi -f $IMAGE_NAME
  fi
done

# 重新构建并启动服务
echo "Rebuilding and starting containers..."
docker-compose -f docker-compose.yml up --build -d