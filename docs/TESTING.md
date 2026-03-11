# 测试指南

## 快速测试

### 1. 准备环境

```bash
# 启动 MySQL
docker run -d --name mysql \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -p 3306:3306 \
  mysql:8.0

# 或使用本地 MySQL
mysql -u root -p
```

### 2. 初始化数据库

```bash
mysql -u root -p123456 < backend/sql/init.sql
```

### 3. 修改配置

编辑 `backend/config/user.yaml`:

```yaml
Database:
  DataSource: root:123456@tcp(localhost:3306)/user_management?charset=utf8mb4&parseTime=true&loc=Local
```

### 4. 启动后端

```bash
cd backend
go mod tidy
go run main.go -f config/user.yaml
```

### 5. 启动前端

```bash
cd frontend
npm install
npm run dev
```

---

## API 测试

### 使用 cURL 测试

#### 1. 用户注册

```bash
curl -X POST http://localhost:8888/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456",
    "email": "test@example.com",
    "nickname": "测试用户"
  }'
```

**响应**:
```json
{
  "user_id": 1,
  "message": "注册成功"
}
```

#### 2. 用户登录

```bash
curl -X POST http://localhost:8888/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456"
  }'
```

**响应**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzEwMzQ1NjAwfQ.xxx",
  "expire": 86400,
  "user": {
    "user_id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "nickname": "测试用户",
    "status": 1,
    "created_at": "2026-03-12 00:00:00"
  }
}
```

**保存 Token**:
```bash
export TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

#### 3. 获取用户列表（需要认证）

```bash
curl -X GET http://localhost:8888/api/v1/users \
  -H "Authorization: Bearer $TOKEN"
```

**响应**:
```json
{
  "total": 2,
  "users": [
    {
      "user_id": 1,
      "username": "testuser",
      "email": "test@example.com",
      "nickname": "测试用户",
      "status": 1,
      "created_at": "2026-03-12 00:00:00"
    }
  ]
}
```

#### 4. 创建用户（需要认证）

```bash
curl -X POST http://localhost:8888/api/v1/users \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "newuser",
    "password": "123456",
    "email": "new@example.com",
    "nickname": "新用户"
  }'
```

#### 5. 获取用户详情（需要认证）

```bash
curl -X GET http://localhost:8888/api/v1/users/1 \
  -H "Authorization: Bearer $TOKEN"
```

#### 6. 更新用户（需要认证）

```bash
curl -X PUT http://localhost:8888/api/v1/users/1 \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "nickname": "新昵称",
    "status": 1
  }'
```

#### 7. 删除用户（需要认证）

```bash
curl -X DELETE http://localhost:8888/api/v1/users/1 \
  -H "Authorization: Bearer $TOKEN"
```

#### 8. 刷新 Token

```bash
curl -X POST http://localhost:8888/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{
    "token": "'$TOKEN'"
  }'
```

#### 9. 用户登出

```bash
curl -X POST http://localhost:8888/api/v1/auth/logout \
  -H "Content-Type: application/json" \
  -d '{
    "token": "'$TOKEN'"
  }'
```

---

## 错误测试

### 1. 未授权访问

```bash
curl -X GET http://localhost:8888/api/v1/users
```

**响应** (401):
```json
{
  "code": 401,
  "message": "missing authorization header"
}
```

### 2. 无效 Token

```bash
curl -X GET http://localhost:8888/api/v1/users \
  -H "Authorization: Bearer invalid_token"
```

**响应** (401):
```json
{
  "code": 401,
  "message": "token has expired"
}
```

### 3. 用户名重复

```bash
curl -X POST http://localhost:8888/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456"
  }'
```

**响应** (400):
```json
{
  "message": "用户名已存在"
}
```

### 4. 密码错误

```bash
curl -X POST http://localhost:8888/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "wrongpassword"
  }'
```

**响应** (400):
```json
{
  "message": "用户名或密码错误"
}
```

---

## 前端测试

### 1. 访问登录页

打开浏览器访问：http://localhost:3000/login

### 2. 测试注册流程

1. 点击"立即注册"
2. 填写注册表单
3. 提交后跳转到登录页

### 3. 测试登录流程

1. 输入用户名和密码
2. 点击登录
3. 成功后跳转到用户列表页

### 4. 测试用户管理

1. **查看列表**: 验证分页和搜索
2. **创建用户**: 点击"新增用户"，填写表单
3. **查看详情**: 点击"详情"按钮
4. **编辑用户**: 点击"编辑"按钮，修改信息
5. **删除用户**: 点击"删除"按钮，确认删除

### 5. 测试登出

1. 点击右上角用户名
2. 选择"退出登录"
3. 验证是否跳转到登录页

### 6. 测试路由守卫

1. 登出状态下访问 http://localhost:3000/
2. 应该自动跳转到登录页
3. 登录后访问受保护页面

---

## 性能测试

### 使用 Apache Bench

```bash
# 测试登录接口
ab -n 1000 -c 10 \
  -H "Content-Type: application/json" \
  -p login.json \
  http://localhost:8888/api/v1/auth/login

# login.json 内容:
# {"username":"admin","password":"123456"}
```

### 使用 wrk

```bash
# 测试用户列表接口
wrk -t12 -c400 -d30s \
  -H "Authorization: Bearer $TOKEN" \
  http://localhost:8888/api/v1/users
```

---

## 常见问题

### Q1: 数据库连接失败

**解决**:
1. 检查 MySQL 是否运行
2. 验证用户名密码
3. 确认数据库已创建

### Q2: Token 过期

**解决**:
1. 使用刷新接口获取新 Token
2. 或重新登录

### Q3: 跨域问题

**解决**:
1. 开发环境已配置代理
2. 生产环境使用 Nginx 反向代理

### Q4: 端口被占用

**解决**:
1. 修改 `user.yaml` 中的端口
2. 或停止占用端口的服务

---

## 测试清单

- [ ] 用户注册成功
- [ ] 用户登录成功
- [ ] Token 正常生成
- [ ] 受保护接口需要认证
- [ ] 无效 Token 被拒绝
- [ ] 用户 CRUD 操作正常
- [ ] 分页功能正常
- [ ] 搜索功能正常
- [ ] 前端路由守卫生效
- [ ] 登出功能正常

---

**最后更新**: 2026-03-12
