根据您提供的代码，以下是`/user/release`接口的文档，涵盖了处理用户发布信息界面的GET, POST, PUT, DELETE请求。

# `/user/release` 接口文档

## GET `/user/release`

此接口用于获取用户发布信息的页面。

### 中间件
+ jwtMiddleware: 验证用户是否已登陆.详细文档查看[这里](../middleware/jwt.md)
+ userIDMatch: 验证用户ID是否匹配.详细文档查看[这里](../middleware/userIDMatch.md)

### 请求

- **方法**：GET
- **URL**：`/user/release`
- **参数**：无

### 响应

#### 成功响应

- **状态码**：`200 OK`
- **内容**：返回`upload.html`页面。

## POST `/user/release`

此接口用于用户发布新的房源信息。支持文件上传。

### 中间件
+ jwtMiddleware: 验证用户是否已登陆.详细文档查看[这里](../middleware/jwt.md)
+ userIDMatch: 验证用户ID是否匹配.详细文档查看[这里](../middleware/userIDMatch.md)

### 请求

- **方法**：POST
- **URL**：`/user/release`
- **表单数据**：需要包含房源的相关信息以及需要上传的文件。文件字段名应为`files`。

### 响应

#### 成功响应

- **状态码**：`200 OK`
- **内容**：

  返回一个JSON对象，包含发布成功的消息和方法类型。

  示例响应体：

  ```json
  {
    "url": "/user/release",
    "method": "POST",
    "message": "发布成功"
  }
  ```

#### 错误响应

- `400 Bad Request`：如果上传文件失败或保存文件失败，将返回错误消息。

  示例错误响应：

  ```json
  {
    "error": "上传文件失败"
  }
  ```

  或

  ```json
  {
    "error": "保存文件失败"
  }
  ```

## PUT `/user/release`

此接口用于更新用户已发布的房源信息。

### 中间件
+ jwtMiddleware: 验证用户是否已登陆.详细文档查看[这里](../middleware/jwt.md)
+ userIDMatch: 验证用户ID是否匹配.详细文档查看[这里](../middleware/userIDMatch.md)

### 请求

- **方法**：PUT
- **URL**：`/user/release`
- **URL参数**：需要在URL中指定`user_id`。

### 响应

#### 成功响应

- **状态码**：`200 OK`
- **内容**：

  返回一个JSON对象，包含请求的URL、方法和用户ID。

  示例响应体：

  ```json
  {
    "message": "/user/release",
    "method": "PUT",
    "user_id": "指定的用户ID"
  }
  ```

## DELETE `/user/release`

此接口用于删除用户发布的房源信息。

### 中间件
+ jwtMiddleware: 验证用户是否已登陆.详细文档查看[这里](../middleware/jwt.md)
+ userIDMatch: 验证用户ID是否匹配.详细文档查看[这里](../middleware/userIDMatch.md)


### 请求

- **方法**：DELETE
- **URL**：`/user/release`
- **URL参数**：需要在URL中指定`user_id`。

### 响应

#### 成功响应

- **状态码**：`200 OK`
- **内容**：

  返回一个JSON对象，包含请求的URL、方法和用户ID。

  示例响应体：

  ```json
  {
    "message": "/user/release",
    "method": "Delete",
    "user_id": "指定的用户ID"
  }
  ```

### 特殊处理

- 对于POST请求，服务器端可能需要处理多个文件的上传。文件保存的目标路径可能包含房源ID、文件名和索引。
- 使用此接口时，需要确保请求中包含有效的用户认证信息，尤其是对于POST, PUT, DELETE方法。

### 注意

- 对于POST方法，确保前端表单设置了`enctype="multipart/form-data"`以支持文件上传。
- 确保服务器配置了足够的存储空间以保存上传的文件。
- 应用适当的权限控制和验证逻辑，确保用户只能编辑或删除自己的房源信息。