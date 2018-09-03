#! /bin/bash

EXEC_PATH=$(cd "$(dirname "$0")"; pwd)
## 项目主目录
PROJECT_DIR=$(dirname ${EXEC_PATH})
## 编译打包后文件
TARGET_DIR=${PROJECT_DIR}/target
## 执行入口
MAIN_CLASS=net.guozhen.ServerLauncher

## 第1步
mvn -f ${PROJECT_DIR} clean install

## 第2步
echo -e "正在10089端口启动netty服务"
java -classpath ${TARGET_DIR}/server.jar net.guozhen.ServerLauncher

## 第3步
# curl http://127.0.0.1:10089/