# douyin-microservice/gateway/Dockerfile

# 使用Go官方镜像作为构建环境
FROM golang:latest as builder

# 设置工作目录
WORKDIR /app

# 使用Alpine镜像作为基础来减小构建的容器大小
FROM alpine:latest


# 设置工作目录
WORKDIR /root/

# 运行gateway应用
CMD ["/bin/sh"]

