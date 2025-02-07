# 使用官方 Go 镜像作为构建阶段
FROM golang:1.22.0 AS builder

# 设置构建阶段的工作目录
WORKDIR /src

# 复制源代码和依赖文件
COPY . .

ENV GOPROXY=https://goproxy.cn

# 下载依赖
RUN go mod download

# 定义构建参数
ARG TIMESTAMP

# 设置编译目标
ENV GOOS=linux
ENV GOARCH=amd64

# 编译应用程序，本地环境下不使用-a加快编译速度
RUN go build -ldflags "-X bootstrap.LastChangedDate=${TIMESTAMP}" -o server_simpletools ./cmd/simpletools/main.go

# 检查编译结果
RUN ls -l /src
RUN ls -l /src/server_simpletools || echo "server_simpletools not found"

# 使用较小的基础镜像来运行应用
FROM alpine:latest

# 安装 tzdata 和 libc 兼容性库
RUN apk add --no-cache tzdata libc6-compat

# 设置时区为上海
ENV TZ=Asia/Shanghai
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# 设置运行阶段的工作目录
WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /src/server_simpletools ./
COPY --from=builder /src/etc ./etc

# 检查复制后的文件
RUN ls -l /app
RUN ls -l ./server_simpletools || echo "server_simpletools not found in /app"

RUN mkdir output
RUN chmod +x ./server_simpletools

# 暴露应用的端口
EXPOSE 1235

# 运行应用程序
CMD ["./server_simpletools"]
