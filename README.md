# 用户管理系统 (User Management System)

基于 go-zero + Vue 3 + MySQL 的用户管理系统

## 🚀 技术栈

- **后端**: Go + go-zero 框架
- **前端**: Vue 3 + TypeScript
- **数据库**: MySQL 8.0
- **认证**: JWT

## 📁 项目结构

```
user-management-system/
├── backend/                 # go-zero 后端
│   ├── api/                # API 定义
│   ├── config/             # 配置文件
│   ├── internal/
│   │   ├── handler/       # 请求处理
│   │   ├── logic/         # 业务逻辑
│   │   ├── model/         # 数据模型
│   │   └── svc/           # 服务上下文
│   └── main.go            # 入口文件
├── frontend/               # Vue 3 前端
└── docs/                   # 项目文档
```

## 🛠️ 快速开始

### 后端启动

```bash
cd backend
go mod tidy
go run main.go -f config/user.yaml
```

### 前端启动

```bash
cd frontend
npm install
npm run dev
```

## 📋 API 接口

### 用户管理

| 接口 | 方法 | 描述 |
|------|------|------|
| `/api/v1/users` | POST | 创建用户 |
| `/api/v1/users/:user_id` | GET | 获取用户详情 |
| `/api/v1/users/:user_id` | PUT | 更新用户 |
| `/api/v1/users/:user_id` | DELETE | 删除用户 |
| `/api/v1/users` | GET | 用户列表 |

## 🔧 配置

编辑 `backend/config/user.yaml`:

```yaml
Database:
  DataSource: root:password@tcp(localhost:3306)/user_management?charset=utf8mb4&parseTime=true&loc=Local

JWT:
  Secret: your-secret-key-change-in-production
```

## 📝 功能规划

### Phase 1 (当前)
- [x] 项目初始化
- [ ] 用户 CRUD 接口
- [ ] 数据库迁移
- [ ] 基础前端页面

### Phase 2
- [ ] JWT 认证
- [ ] 登录/注册
- [ ] 权限控制

### Phase 3
- [ ] 角色管理
- [ ] 操作日志
- [ ] 数据导出

## 📄 License

MIT
