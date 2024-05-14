# 在线房屋交易平台

## 配置文件
- 项目根目录下的`config.json`文件中配置了数据库连接信息，如需修改请在此文件中修改。

## 项目运行
在项目根目录下执行以下命令：
```shell
go run main.go
```

## 项目接口文档
可以直接查看[接口文档](./docs/swagger.json)
或在项目启动后访问[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)查看。

项目默认监听`8080`端口，可在`config.json`中修改.