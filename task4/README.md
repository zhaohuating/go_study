# task4 项目说明

## 项目简介
本项目是一个基于 Go 语言开发的博客系统后端，包含用户注册登录、文章管理、评论管理等功能，使用 Gin 框架搭建 HTTP 服务，GORM 作为 ORM 工具操作数据库，支持 JWT 身份认证。

## 运行环境
- Go 1.24 及以上版本
- MySQL 数据库（默认配置）或 SQLite 数据库(未做测试)
- Git（用于克隆代码）

## 依赖安装步骤

1. **克隆项目代码**
   ```bash
   git clone https://github.com/zhaohuating/go_study.git
   cd go_study/task4
   ```

2. **初始化模块并安装依赖**
   项目已包含 `go.mod` 和 `go.sum` 文件，执行以下命令安装依赖：
   ```bash
   go mod tidy
   ```
   该命令会自动下载并安装项目所需的所有依赖包，包括：
    - Gin 框架：用于构建 HTTP 服务
    - GORM：ORM 框架，用于数据库操作
    - JWT 相关库：用于身份认证
    - YAML 解析库：用于解析配置文件
    - 数据库驱动：支持 MySQL 和 SQLite

## 配置文件修改
项目配置文件位于 `config/config.yaml`，可根据实际环境修改以下关键配置：

1. **服务器配置**
   ```yaml
   server:
     port: 8082  # 服务监听端口
   ```

2. **数据库配置**
   ```yaml
   database:
     driver: "mysql"  # 数据库驱动，可选 mysql 或 sqlite
     dsn: "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"  # 数据库连接字符串
     max_open_conns: 100  # 最大打开连接数
     max_idle_conns: 20   # 最大空闲连接数
     conn_max_lifetime: 3600  # 连接最大存活时间(秒)
   ```

3. **JWT 配置**
   ```yaml
   jwt:
     secret: "myjwt@secretkey_"  # JWT 签名密钥
     expire_hours: 24  # 过期时间(小时)
   ```
4. **日志 配置**
   ```yaml
   log:
     level: "debug" # panic fatal error warn/warning info debug trace
     max_age: 7  # 日志保存时间 （天）
     output_path: "./logs/logs-"
   ```

## 启动方式

1. **直接启动**
   在项目根目录（`task4` 目录）下执行：
   ```bash
   go run main.go
   ```
   服务会根据配置文件中的端口启动，默认地址为 `http://localhost:8082`

2. **构建可执行文件后启动**
   ```bash
   # 构建
   go build -o blog_server main.go
   # 启动
   ./blog_server  # Linux/Mac
   # 或
   blog_server.exe  # Windows
   ```

## 主要功能接口
- 公开接口（无需认证）
    - `POST /api/register`：用户注册
    - `POST /api/login`：用户登录

- 需认证接口（需在请求头携带 JWT Token）
    - 文章管理：`POST /api/post`（创建）、`GET /api/post`（列表）、`GET /api/post/:id`（详情）、`PUT /api/post/:id`（更新）、`DELETE /api/post/:id`（删除）
    - 评论管理：`POST /api/comment`（创建）、`GET /api/comment/:postID`（获取某篇文章的评论列表）

---

---

# API

Base URLs:

# Authentication

# Default

## POST 注册

POST /register

> Body 请求参数

```json
{
  "username": "韩梅梅",
  "password": "123456",
  "email": "hanmeimei@qq.com"
}
```

### 请求参数

| 名称 | 位置 | 类型   | 必选 | 说明 |
| ---- | ---- | ------ | ---- | ---- |
| body | body | object | 是   | none |

> 返回示例

> 200 Response

```json
{
  "message": "User registered successfully"
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

## POST 登录

POST /login

> Body 请求参数

```json
{
  "username": "韩梅梅",
  "password": "123456"
}
```

### 请求参数

| 名称 | 位置 | 类型   | 必选 | 说明 |
| ---- | ---- | ------ | ---- | ---- |
| body | body | object | 是   | none |

> 返回示例

> 200 Response

```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTU2MTEyMDIsImlkIjoyLCJ1c2VybmFtZSI6IumfqeaiheaihSJ9.AnGPujal8GqnqoDL-6EhDfbioR60H8JOZXYn1eLdMXc",
  "user": {
    "id": 2,
    "username": "韩梅梅"
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

## POST 添加文章

POST /user/addpost

> Body 请求参数

```json
{
  "title": "这是韩梅梅的第一篇博客",
  "content": "第一次写博客真的好开心，希望大家能喜欢"
}
```

### 请求参数

| 名称          | 位置   | 类型   | 必选 | 说明 |
| ------------- | ------ | ------ | ---- | ---- |
| Authorization | header | string | 是   | none |
| body          | body   | object | 是   | none |

> 返回示例

> 200 Response

```json
{
  "message": "Post created successfully"
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

## GET 获取文章列表带分页

GET /post

### 请求参数

| 名称          | 位置   | 类型    | 必选 | 说明 |
| ------------- | ------ | ------- | ---- | ---- |
| page          | query  | integer | 是   | none |
| pageSize      | query  | integer | 是   | none |
| Authorization | header | string  | 是   | none |

> 返回示例

```json
{
  "page": 1,
  "pageSize": 10,
  "total": 1,
  "totalPage": 1,
  "items": [
    {
      "ID": 1,
      "CreatedAt": "2025-08-18T09:54:27.024+08:00",
      "UpdatedAt": "2025-08-18T09:54:27.024+08:00",
      "DeletedAt": null,
      "Title": "这是韩梅梅的第一篇博客",
      "Content": "第一次写博客真的号开心，希望大家能喜欢",
      "UserID": 2
    }
  ]
}
```

```json
{
  "code": 500,
  "message": "查询失败"
}
```

```json
{
  "page": 1,
  "pageSize": 10,
  "total": 1,
  "totalPage": 1,
  "items": [
    {
      "ID": 2,
      "CreatedAt": "2025-08-18T17:19:28.175+08:00",
      "UpdatedAt": "2025-08-18T17:19:28.175+08:00",
      "DeletedAt": null,
      "Title": "这是韩梅梅的第一篇博客",
      "Content": "",
      "UserID": 0
    }
  ]
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

## PUT 更新博客

PUT /post/1

/post/{postid}.

postid：文章ID

> Body 请求参数

```json
{
  "title": "这是韩梅梅的第一篇博客",
  "content": "第一次写博客真的好开心，希望大家能喜欢，感谢大家"
}
```

### 请求参数

| 名称          | 位置   | 类型   | 必选 | 说明 |
| ------------- | ------ | ------ | ---- | ---- |
| Authorization | header | string | 是   | none |
| body          | body   | object | 是   | none |

> 返回示例

> 200 Response

```json
{
  "message": "Post updated successfully"
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

## DELETE 删除博客

DELETE /post/1

/post/{postid}

postid：文章ID

### 请求参数

| 名称          | 位置   | 类型   | 必选 | 说明 |
| ------------- | ------ | ------ | ---- | ---- |
| Authorization | header | string | 是   | none |

> 返回示例

> 200 Response

```json
{
  "message": "Post deleted successfully"
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称      | 类型   | 必选 | 约束 | 中文名 | 说明 |
| --------- | ------ | ---- | ---- | ------ | ---- |
| » message | string | true | none |        | none |

## POST 新增评论

POST /comment

> Body 请求参数

```json
{
  "content": "博主写的太好了，支持博主",
  "postID": 2
}
```

### 请求参数

| 名称          | 位置   | 类型   | 必选 | 说明 |
| ------------- | ------ | ------ | ---- | ---- |
| Authorization | header | string | 是   | none |
| body          | body   | object | 是   | none |

> 返回示例

> 200 Response

```json
{
  "message": "Comment created successfully"
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

## GET 获取文章的评论

GET /comment/2

### 请求参数

| 名称          | 位置   | 类型    | 必选 | 说明     |
| ------------- | ------ | ------- | ---- | -------- |
| page          | query  | integer | 是   | 页数     |
| pageSize      | query  | integer | 是   | 每页条数 |
| Authorization | header | string  | 是   | none     |

> 返回示例

> 200 Response

```json
{
  "page": 1,
  "pageSize": 20,
  "total": 2,
  "totalPage": 1,
  "items": [
    {
      "ID": 1,
      "CreatedAt": "2025-08-18T17:20:10.816+08:00",
      "UpdatedAt": "2025-08-18T17:20:10.816+08:00",
      "DeletedAt": null,
      "Content": "博主写的太好了，支持博主",
      "UserID": 2,
      "PostID": 2
    },
    {
      "ID": 2,
      "CreatedAt": "2025-08-18T18:12:03+08:00",
      "UpdatedAt": "2025-08-18T18:12:16+08:00",
      "DeletedAt": null,
      "Content": "支持博主",
      "UserID": 1,
      "PostID": 2
    }
  ]
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

## GET 获取文章详细信息

GET /post/2

/post/{postid}

postid：文章ID

### 请求参数

| 名称          | 位置   | 类型   | 必选 | 说明 |
| ------------- | ------ | ------ | ---- | ---- |
| Authorization | header | string | 是   | none |

> 返回示例

> 200 Response

```json
{
  "ID": 2,
  "CreatedAt": "2025-08-18T17:19:28.175+08:00",
  "UpdatedAt": "2025-08-18T17:19:28.175+08:00",
  "DeletedAt": null,
  "Title": "这是韩梅梅的第一篇博客",
  "Content": "第一次写博客真的好开心，希望大家能喜欢",
  "UserID": 2
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 