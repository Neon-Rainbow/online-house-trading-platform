# 在线房屋交易平台

本项目为苏州大学 综合项目实践 SOEN2014 课程作业

本人负责完成项目的后端部分

马后炮的说,这个项目其实有点简陋.双token刷新是项目ddl一周之前完成的,项目本来初始设计时是有复杂的权限校验模块以及在线聊天模块的,可惜这两个在项目中都为能实现.其中权限校验模块中,由于管理员的一些接口是在项目ddl前一个月的时候做完的,在线聊天功能当时不会websocket,也没搞出来.可能因为这两个模块未完成导致最终这个项目的得分并不够高.在暑假时,我重新设计了权限校验模块以及使用websocket的在线交流模块,项目地址分别如下:

+   权限校验模块 [Github](https://github.com/Neon-Rainbow/JWT_authorization) [Gitee](https://gitee.com/Aspirin-Byte/JWT_authorization)
+   在线聊天模块 [GitHub](https://github.com/Neon-Rainbow/Gin_Websocket_Chat) [Gitee](https://gitee.com/Aspirin-Byte/Gin_Websocket_Chat)

其中权限校验模块提供了完整的HTTP接口以及gRPC接口

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
