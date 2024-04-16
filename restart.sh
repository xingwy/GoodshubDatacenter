#!/bin/bash

# 编译 main.go 文件
echo "go build"
go mod tidy
go build -o service_app main.go

# 确保 ./bin 目录存在
mkdir -p ./bin

# 移动编译后的可执行文件到 ./bin 目录
mv -f ./service_app ./bin/

# 赋权限
chmod +x ./bin/service_app

# 进程清除
# killall service_app # kill service_app service_app


./bin/service_app
