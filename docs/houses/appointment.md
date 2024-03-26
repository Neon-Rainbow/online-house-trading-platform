# `/houses/appointment` 接口文档

## POST `/houses/appointment`

该接口用于实现用户对房屋的预约。用户需要在请求头部添加有效的JWT令牌才能正常访问此接口。

### 鉴权

- 需要在请求头中携带有效的JWT令牌，格式为`Authorization: Bearer <token>`或`Authorization: <token>`。

### 请求

需要在请求体中提供以下JSON格式的数据：

```json
{
  "house_id": "房屋ID",
  "appointment_date": "预约日期"
}
```

### 响应

#### 成功响应

- `200 OK`：预约成功。返回JSON格式的数据，包含预约成功的消息。

  响应体示例：

  ```json
  {
    "message": "预约成功",
    "url": "/houses/appointment"
  }
  ```

#### 错误响应

- `400 Bad Request`：请求参数错误。如果是因为参数绑定失败或其他请求参数问题，将返回具体的错误信息。

	示例:
	```json
	{
	  "error":  "请求参数错误",
	  "detail": err
	}
	```


- `401 Unauthorized`：请求未携带token或token无效。如果未提供token或提供的token无法验证，将返回错误信息。

  示例：

  ```json
  {
    "error": "请求未携带token，无权限访问",
    "detail": err
  }
  ```

  或

  ```json
  {
    "error": "无效的token",
    "detail": err
  }
  ```

- `500 Internal Server Error`：无法获取数据库连接或无法获取用户ID。这可能是服务器内部配置错误或数据库服务不可用。

### 特殊处理

- 用户必须通过JWT令牌验证才能进行房屋预约操作。JWT令牌应在用户登录成功后获得，并在此接口的请求头中提供。
- 预约信息包括房屋ID和预约日期，这些信息将被存入数据库。成功预约后，会返回预约成功的消息。

### 注意

请确保使用HTTPS协议来保护用户的敏感信息安全，特别是在携带JWT令牌进行认证的请求中。
