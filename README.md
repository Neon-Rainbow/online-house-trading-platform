# 在线房屋交易平台

## git commit 规范

在本地修改代码时,禁止直接修改main分支.请务必使用git branch创建新的分支,并在新分支上进行修改.
修改完后,请使用git merge将新分支合并到main分支.

commit message格式
```
<type>(<scope>): <subject>
```
### type(必须)

用于说明git commit的类别，只允许使用下面的标识。

feat：新功能（feature）。

fix/to：修复bug，可以是QA发现的BUG，也可以是研发自己发现的BUG。

fix：产生diff并自动修复此问题。适合于一次提交直接修复问题
to：只产生diff不自动修复此问题。适合于多次提交。最终修复问题提交时使用fix
docs：文档（documentation）。

style：格式（不影响代码运行的变动）。

refactor：重构（即不是新增功能，也不是修改bug的代码变动）。

perf：优化相关，比如提升性能、体验。

test：增加测试。

chore：构建过程或辅助工具的变动。

revert：回滚到上一个版本。

merge：代码合并。

sync：同步主线或分支的Bug。

### scope(可选)

scope用于说明 commit 影响的范围，比如数据层、控制层、视图层等等，视项目不同而不同。

例如在Angular，可以是location，browser，compile，compile，rootScope， ngHref，ngClick，ngView等。如果你的修改影响了不止一个scope，你可以使用*代替。

### subject(必须)

subject是commit目的的简短描述，不超过50个字符。



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