# 在线房屋交易平台

# Online House Trading Platform

![Top Language](https://img.shields.io/github/languages/top/Neon-Rainbow/online-house-trading-platform)
![Language Count](https://img.shields.io/github/languages/count/Neon-Rainbow/online-house-trading-platform)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub contributors](https://img.shields.io/github/contributors/Neon-Rainbow/online-house-trading-platform)](https://github.com/Neon-Rainbow/online-house-trading-platform/graphs/contributors)


## Docker

使用下列指令构建Docker镜像:
```bash
docker build -t online-house-trading-platform .
```

使用下列指令运行Docker镜像:
```bash
docker run -d --name online-house-trading-platform -p 8080:8080 online-house-trading-platform 
```

或者使用Docker Compose:
```bash
docker compose up -d
```
可以使用
```bash
docker compose down
```
移除容器

注意:mysql未在Docker中配置

## 配置文件样例
```json
{
  "database": {
    "host": "",
    "port": 3306,
    "user": "",
    "password": "",
    "dbname": ""
  },
  "redis": {
    "host": "",
    "port": 6379,
    "password": "",
    "db": 0
  },
  "jwtSecret": "",
  "passwordSecret": "",
  "logFilePath": "./application.log",
  "port": 8080,
  "ginMode": "debug",
  "zapLogLever": "debug",
  "admin_register_secret_key": ""
}
```



## 配置文件
- 项目根目录下的`config.json`文件中配置了数据库连接信息，如需修改请在此文件中修改。

## 项目运行
在项目根目录下执行以下命令：
```shell
go run main.go
```

## 项目接口文档
可以直接查看[接口文档](项目文档/接口文档/房屋交易平台.postman_collection.json),需要将该接口文档导入Postman使用

项目默认监听`8080`端口，可在`config.json`中修改.

## 项目日志文件

项目的日志文件会在[application.log](./application.log)中.如果需要格式化 JSON 数据,可以运行项目根目录下的[format_log.py](./format_log.py)

可以在项目根目录下执行以下指令:
```shell
python format_log.py
```
