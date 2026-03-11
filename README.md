# 用户管理系统 (User Management System)

基于 go-zero + Vue 3 + MySQL 的用户管理系统

[![GitHub](https://img.shields.io/github/license/ZHUCHANG1999/user-management-system)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.21-blue)](https://golang.org)
[![Vue](https://img.shields.io/badge/Vue-3.4-green)](https://vuejs.org)

## 🚀 技术栈

- **后端**: Go 1.21 + go-zero v1.6
- **前端**: Vue 3 + Vite + Element Plus
- **数据库**: MySQL 8.0
- **认证**: JWT (待实现)

## 📁 项目结构

```
user-management-system/
├── backend/                 # go-zero 后端
│   ├── api/                # API 定义 (.api 文件)
│   ├── config/             # 配置文件
│   ├── internal/
│   │   ├── handler/       # HTTP 处理器
│   │   ├── logic/         # 业务逻辑
│   │   ├── model/         # 数据模型
│   │   └── svc/           # 服务上下文
│   ├── sql/               # 数据库脚本
│   └── main.go            # 入口文件
├── frontend/               # Vue 3 前端
│   ├── src/
│   │   ├── api/           # API 调用
│   │   ├── router/        # 路由配置
│   │   ├── views/         # 页面组件
│   │   ├── App.vue        # 根组件
│   │   └── main.js        # 入口文件
│   ├── index.html
│   └── package.json
├── docs/                   # 项目文档
└── README.md
```

## 🛠️ 快速开始

### 环境要求

- Go 1.21+
- Node.js 18+
- MySQL 8.0+

### 1. 数据库初始化

```bash
# 登录 MySQL
mysql -u root -p

# 执行初始化脚本
source backend/sql/init.sql
```

### 2. 后端启动

```bash
cd backend

# 下载依赖
go mod tidy

# 修改配置文件 (config/user.yaml)
# 更新数据库连接信息

# 运行服务
go run main.go -f config/user.yaml
```

服务启动后访问：http://localhost:8888

### 3. 前端启动

```bash
cd frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

前端访问：http://localhost:3000

## 📋 API 接口

### 用户管理

| 接口 | 方法 | 描述 | 请求参数 |
|------|------|------|----------|
| `/api/v1/users` | POST | 创建用户 | username, password, email, nickname |
| `/api/v1/users/:user_id` | GET | 获取用户详情 | user_id (path) |
| `/api/v1/users/:user_id` | PUT | 更新用户 | user_id (path), email, nickname, status |
| `/api/v1/users/:user_id` | DELETE | 删除用户 | user_id (path) |
| `/api/v1/users` | GET | 用户列表 | page, page_size, username (query) |

### 请求示例

**创建用户**
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

**获取用户列表**
```bash
curl http://localhost:8888/api/v1/users?page=1&page_size=10
```

## 🔧 配置说明

### 后端配置 (backend/config/user.yaml)

```yaml
Database:
  DataSource: root:your_password@tcp(localhost:3306)/user_management?charset=utf8mb4&parseTime=true&loc=Local

JWT:
  Secret: your-secret-key-change-in-production
  Expire: 86400
```

### 前端代理 (frontend/vite.config.js)

已配置 API 代理，开发环境下 `/api` 请求自动转发到后端 `http://localhost:8888`

## 📝 功能清单

### Phase 1 ✅ (已完成)
- [x] 项目初始化
- [x] 用户 CRUD 接口
- [x] 数据库迁移脚本
- [x] 前端用户列表页面
- [x] 前端用户创建/编辑/详情页面
- [x] GitHub 仓库搭建

### Phase 2 ✅ (已完成)
- [x] JWT 认证
- [x] 登录/注册页面
- [x] 密码加密存储 (bcrypt)
- [x] 会话管理
- [x] Token 刷新机制
- [x] 路由守卫
- [x] 401 自动跳转

### Phase 3 ✅ (已完成)
- [x] 角色管理 (CRUD)
- [x] 权限管理 (CRUD)
- [x] 角色 - 权限关联
- [x] 用户 - 角色关联
- [x] RBAC 权限验证
- [x] 默认角色和权限初始化
- [ ] 操作日志
- [ ] 数据导出

## 🎨 界面预览

### 用户列表
- 支持分页
- 支持按用户名搜索
- 支持状态显示

### 用户管理
- 创建用户（带表单验证）
- 编辑用户信息
- 查看用户详情
- 删除用户（带确认）

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 License

MIT License

## 📞 联系

- Author: ZHUCHANG1999
- GitHub: https://github.com/ZHUCHANG1999
