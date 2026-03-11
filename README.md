# 用户管理系统

基于 go-zero + Vue 3 的用户管理系统。

## 技术栈

**后端：**
- Go 1.19+
- go-zero 微服务框架
- MySQL 8.0+
- JWT 鉴权

**前端：**
- Vue 3 + TypeScript
- Vite 5
- Element Plus UI
- Pinia 状态管理
- Vue Router 4
- Axios

## 项目结构

```
user-management-system/
├── backend/                 # 后端代码
│   ├── etc/                # 配置文件
│   ├── internal/           # 内部包
│   │   ├── config/        # 配置结构
│   │   ├── handler/       # HTTP 处理器
│   │   ├── logic/         # 业务逻辑
│   │   └── svc/           # 服务上下文
│   ├── types/             # 类型定义
│   ├── main.go            # 入口文件
│   ├── user.api           # API 定义
│   └── schema.sql         # 数据库脚本
├── frontend/              # 前端代码
│   ├── src/
│   │   ├── api/          # API 接口
│   │   ├── assets/       # 静态资源
│   │   ├── components/   # 组件
│   │   ├── router/       # 路由配置
│   │   ├── stores/       # Pinia 状态
│   │   ├── views/        # 页面视图
│   │   ├── App.vue       # 根组件
│   │   └── main.ts       # 入口文件
│   ├── index.html
│   ├── package.json
│   ├── tsconfig.json
│   └── vite.config.ts
└── README.md
```

## 快速开始

### 1. 数据库初始化

```bash
# 创建数据库并导入表结构
mysql -u root -p < backend/schema.sql
```

### 2. 后端启动

```bash
cd backend

# 安装依赖
go mod tidy

# 修改配置文件 etc/user-api.yaml 中的数据库连接信息

# 运行
go run main.go -f etc/user-api.yaml
```

后端服务将在 `http://localhost:8888` 启动

### 3. 前端启动

```bash
cd frontend

# 安装依赖
npm install

# 开发模式运行
npm run dev
```

前端服务将在 `http://localhost:3000` 启动

## API 接口

| 接口 | 方法 | 描述 | 鉴权 |
|------|------|------|------|
| /api/user/register | POST | 用户注册 | ❌ |
| /api/user/login | POST | 用户登录 | ❌ |
| /api/user/info | GET | 获取用户信息 | ✅ |
| /api/user/update | POST | 更新用户信息 | ✅ |
| /api/user/list | GET | 获取用户列表 | ✅ |
| /api/user/delete | POST | 删除用户 | ✅ (仅管理员) |

## 功能特性

- ✅ 用户注册/登录
- ✅ JWT Token 鉴权
- ✅ 用户资料管理（CRUD）
- ✅ 角色权限控制（管理员/普通用户）
- ✅ 用户列表分页展示
- ✅ 软删除机制
- ✅ 前端路由守卫
- ✅ 响应式 UI 设计

## 默认管理员

首次使用需要手动在数据库创建管理员账号：

```sql
-- 密码：admin123 (MD5 加密)
INSERT INTO `user` (username, password_hash, email, nickname, role, status) 
VALUES ('admin', '0192023a7bb5ebf883036ea189e51293', 'admin@example.com', '管理员', 'admin', 1);
```

## 开发说明

### 后端

- 遵循 go-zero 标准项目结构
- Handler → Logic → Model 三层架构
- 使用 sqlx 进行数据库操作
- JWT Token 有效期 24 小时（可在配置中调整）

### 前端

- TypeScript 严格模式
- Composition API + `<script setup>`
- Element Plus 组件库
- Axios 拦截器统一处理鉴权和错误

## 下一步扩展

- [ ] 邮箱验证
- [ ] 密码重置
- [ ] 第三方登录（微信/GitHub）
- [ ] 细粒度权限控制（RBAC）
- [ ] 操作日志
- [ ] 文件上传（头像）
- [ ] Docker 部署

## License

MIT
