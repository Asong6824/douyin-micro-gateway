#!/bin/bash

# 克隆etcd的v3.5.11版本
git clone -b v3.5.11 https://github.com/etcd-io/etcd.git

# 进入克隆的etcd目录
cd etcd

# 执行构建脚本
./build.sh

# 将etcd二进制文件的路径添加到PATH环境变量中
export PATH="$PATH:`pwd`/bin"
