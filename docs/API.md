# API 接口文档

## 基础信息

- **Base URL**: `http://localhost:8888/api/v1`
- **数据格式**: JSON
- **字符编码**: UTF-8
- **认证方式**: JWT Bearer Token（部分接口需要）

---

## 认证接口（公开）

### 1. 用户登录

**接口**: `POST /auth/login`

**请求参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 用户名 |
| password | string | 是 | 密码 |

**请求示例**:
```json
{
  "username": "admin",
  "password": "123456"
}
```

**响应示例**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expire": 86400,
  "user": {
    "user_id": 1,
    "username": "admin",
    "email": "admin@example.com",
    "nickname": "管理员",
    "status": 1,
    "created_at": "2026-03-12 00:00:00"
  }
}
```

---

### 2. 用户注册

**接口**: `POST /auth/register`

**请求参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 用户名（3-20 字符） |
| password | string | 是 | 密码（最少 6 位） |
| email | string | 否 | 邮箱 |
| nickname | string | 否 | 昵称 |

**响应示例**:
```json
{
  "user_id": 3,
  "message": "注册成功"
}
```

---

### 3. 刷新 Token

**接口**: `POST /auth/refresh`

**请求参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| token | string | 是 | 当前 Token |

**响应示例**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expire": 86400
}
```

---

### 4. 用户登出

**接口**: `POST /auth/logout`

**请求参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| token | string | 是 | 当前 Token |

**响应示例**:
```json
{
  "message": "登出成功"
}
```

---

## 用户管理接口（需要认证）

以下接口需要在请求头中携带 JWT Token：

```
Authorization: Bearer <your_token>
```

### 5. 创建用户

**接口**: `POST /users`

**请求参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| username | string | 是 | 用户名（3-20 字符） |
| password | string | 是 | 密码（最少 6 位） |
| email | string | 否 | 邮箱地址 |
| nickname | string | 否 | 昵称 |

**请求示例**:
```json
{
  "username": "testuser",
  "password": "123456",
  "email": "test@example.com",
  "nickname": "测试用户"
}
```

**响应示例**:
```json
{
  "user_id": 1,
  "message": "用户创建成功"
}
```

**错误响应**:
```json
{
  "code": 400,
  "message": "用户名已存在"
}
```

---

### 2. 获取用户详情

**接口**: `GET /users/:user_id`

**路径参数**:

| 参数 | 类型 | 说明 |
|------|------|------|
| user_id | int64 | 用户 ID |

**响应示例**:
```json
{
  "user_id": 1,
  "username": "admin",
  "email": "admin@example.com",
  "nickname": "管理员",
  "status": 1,
  "created_at": "2026-03-12 00:00:00"
}
```

---

### 3. 更新用户

**接口**: `PUT /users/:user_id`

**路径参数**:

| 参数 | 类型 | 说明 |
|------|------|------|
| user_id | int64 | 用户 ID |

**请求参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| email | string | 否 | 邮箱地址 |
| nickname | string | 否 | 昵称 |
| status | int | 否 | 状态（1:正常，0:禁用） |

**请求示例**:
```json
{
  "email": "newemail@example.com",
  "nickname": "新昵称",
  "status": 1
}
```

**响应示例**:
```json
{
  "message": "用户更新成功"
}
```

---

### 4. 删除用户

**接口**: `DELETE /users/:user_id`

**路径参数**:

| 参数 | 类型 | 说明 |
|------|------|------|
| user_id | int64 | 用户 ID |

**响应示例**:
```json
{
  "message": "用户删除成功"
}
```

**说明**: 使用软删除，数据不会从数据库物理删除。

---

### 5. 获取用户列表

**接口**: `GET /users`

**查询参数**:

| 参数 | 类型 | 必填 | 默认值 | 说明 |
|------|------|------|--------|------|
| page | int | 否 | 1 | 页码 |
| page_size | int | 否 | 10 | 每页数量 |
| username | string | 否 | - | 用户名模糊搜索 |

**请求示例**:
```
GET /users?page=1&page_size=10&username=admin
```

**响应示例**:
```json
{
  "total": 2,
  "users": [
    {
      "user_id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "nickname": "管理员",
      "status": 1,
      "created_at": "2026-03-12 00:00:00"
    },
    {
      "user_id": 2,
      "username": "user1",
      "email": "user1@example.com",
      "nickname": "用户 1",
      "status": 1,
      "created_at": "2026-03-12 00:01:00"
    }
  ]
}
```

---

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

---

## cURL 测试示例

### 创建用户
```bash
curl -X POST http://localhost:8888/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456",
    "email": "test@example.com",
    "nickname": "测试用户"
  }'
```

### 获取用户列表
```bash
curl http://localhost:8888/api/v1/users?page=1&page_size=10
```

### 获取用户详情
```bash
curl http://localhost:8888/api/v1/users/1
```

### 更新用户
```bash
curl -X PUT http://localhost:8888/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "nickname": "新昵称",
    "status": 1
  }'
```

### 删除用户
```bash
curl -X DELETE http://localhost:8888/api/v1/users/1
```

---

## Postman 集合

可导入以下 Postman 集合进行测试：

```json
{
  "info": {
    "name": "User Management API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Create User",
      "request": {
        "method": "POST",
        "header": [{"key": "Content-Type", "value": "application/json"}],
        "url": "http://localhost:8888/api/v1/users",
        "body": {
          "mode": "raw",
          "raw": "{\n  \"username\": \"testuser\",\n  \"password\": \"123456\",\n  \"email\": \"test@example.com\",\n  \"nickname\": \"测试用户\"\n}"
        }
      }
    },
    {
      "name": "Get Users",
      "request": {
        "method": "GET",
        "url": "http://localhost:8888/api/v1/users?page=1&page_size=10"
      }
    }
  ]
}
```
