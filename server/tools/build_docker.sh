#!/bin/bash

# 设置镜像名称
IMAGE_NAME="ttptl-admin-server"

# 移动到父目录
cd ..

# 获取当前日期和时间
datetime=$(date +"%Y%m%d%H%M%S")

# 提取年、月、日、时、分、秒
year=${datetime:0:4}
month=${datetime:4:2}
day=${datetime:6:2}
hour=${datetime:8:2}
minute=${datetime:10:2}
second=${datetime:12:2}

# 生成时间戳
timestamp="$year$month$day$hour$minute$second"

# 构建 Docker 镜像
docker build -f ./Dockerfile --build-arg TIMESTAMP="$timestamp" -t "$IMAGE_NAME" .

# 检查构建是否成功
if [ $? -ne 0 ]; then
    echo "Docker build failed!"
    exit 1
fi

# 重启更新后的镜像
CONTAINER_NAME="ttptl-admin-server"
PORT="8081"

# 获取容器信息
CREATED=$(docker inspect --format="{{.Created}}" "$CONTAINER_NAME")

# 获取旧容器的镜像ID
OLD_IMAGE_ID=$(docker inspect --format="{{.Image}}" "$CONTAINER_NAME")

# 打印创建时间和版本号
echo "old container $CONTAINER_NAME was created at: $CREATED"

# 停止并删除旧容器
echo "stop old container:"
docker stop "$CONTAINER_NAME" || echo "No container to stop."
echo "remove old container:"
docker rm "$CONTAINER_NAME" || echo "No container to remove."

# 为旧镜像添加新的标签
echo "Tagging old image with timestamp:$timestamp"
docker tag "$OLD_IMAGE_ID" "$IMAGE_NAME:$timestamp"

# 启动新的容器
echo "start new container $CONTAINER_NAME:"
docker run -d --name "$CONTAINER_NAME" -p "$PORT":"$PORT" "$IMAGE_NAME"
echo "new container $CONTAINER_NAME started on port $PORT."
