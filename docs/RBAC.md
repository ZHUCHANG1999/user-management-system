# RBAC 权限管理指南

## 概述

本系统采用 **RBAC (Role-Based Access Control)** 权限模型，通过角色作为中介，将用户与权限关联。

```
用户 (User) ──→ 角色 (Role) ──→ 权限 (Permission)
```

## 核心概念

### 1. 用户 (User)
系统的使用者，可以拥有一个或多个角色。

### 2. 角色 (Role)
权限的集合，代表一类用户的权限范围。

**默认角色**:
- **超级管理员 (super_admin)**: 拥有所有权限
- **管理员 (admin)**: 拥有除权限管理外的所有权限
- **普通用户 (user)**: 基础查看权限
- **访客 (guest)**: 只读权限

### 3. 权限 (Permission)
对系统资源的访问和操作权利。

**权限类型**:
- **menu (菜单)**: 导航菜单权限
- **button (按钮)**: 页面按钮权限
- **api (接口)**: API 接口调用权限

## 数据库设计

### 表结构

```sql
-- 用户表
users (user_id, username, password, email, ...)

-- 角色表
roles (role_id, role_name, role_code, description, status, ...)

-- 权限表
permissions (permission_id, perm_name, perm_code, perm_type, resource, action, ...)

-- 用户 - 角色关联表
user_roles (user_id, role_id)

-- 角色 - 权限关联表
role_permissions (role_id, permission_id)
```

### ER 图

```
┌─────────────┐       ┌──────────────┐       ┌─────────────┐
│    users    │       │  user_roles  │       │    roles    │
├─────────────┤       ├──────────────┤       ├─────────────┤
│ user_id (PK)│◄──────│ user_id (FK) │       │ role_id (PK)│
│ username    │       │ role_id (FK) │──────►│ role_name   │
│ password    │       └──────────────┘       │ role_code   │
└─────────────┘                              └──────┬──────┘
                                                    │
                                                    │
                                           ┌────────▼────────┐
                                           │role_permissions │
                                           ├─────────────────┤
                                           │ role_id (FK)    │
                                           │ permission_id(FK)│
                                           └────────┬────────┘
                                                    │
                                                    │
                                           ┌────────▼────────┐
                                           │  permissions    │
                                           ├─────────────────┤
                                           │permission_id(PK)│
                                           │ perm_name       │
                                           │ perm_code       │
                                           │ perm_type       │
                                           └─────────────────┘
```

## API 接口

### 角色管理

#### 1. 创建角色
```http
POST /api/v1/roles
Authorization: Bearer <token>

{
  "role_name": "经理",
  "role_code": "manager",
  "description": "部门经理角色"
}
```

#### 2. 获取角色详情
```http
GET /api/v1/roles/:role_id
Authorization: Bearer <token>
```

#### 3. 更新角色
```http
PUT /api/v1/roles/:role_id
Authorization: Bearer <token>

{
  "role_name": "高级经理",
  "description": "高级部门经理",
  "status": 1
}
```

#### 4. 删除角色
```http
DELETE /api/v1/roles/:role_id
Authorization: Bearer <token>
```

#### 5. 角色列表
```http
GET /api/v1/roles?page=1&page_size=10&role_name=经理
Authorization: Bearer <token>
```

### 权限管理

#### 1. 创建权限
```http
POST /api/v1/permissions
Authorization: Bearer <token>

{
  "perm_name": "导出数据",
  "perm_code": "user:export",
  "perm_type": "button",
  "resource": "/api/v1/users/export",
  "action": "POST",
  "description": "导出用户数据"
}
```

#### 2. 权限列表
```http
GET /api/v1/permissions?page=1&page_size=10&perm_type=api
Authorization: Bearer <token>
```

### 角色 - 权限关联

#### 1. 为角色分配权限
```http
POST /api/v1/roles/:role_id/permissions
Authorization: Bearer <token>

{
  "permission_ids": [1, 2, 3, 4, 5]
}
```

#### 2. 获取角色的权限
```http
GET /api/v1/roles/:role_id/permissions
Authorization: Bearer <token>
```

### 用户 - 角色关联

#### 1. 为用户分配角色
```http
POST /api/v1/users/:user_id/roles
Authorization: Bearer <token>

{
  "role_ids": [1, 2]
}
```

#### 2. 获取用户的角色
```http
GET /api/v1/users/:user_id/roles
Authorization: Bearer <token>
```

#### 3. 获取用户的权限
```http
GET /api/v1/users/:user_id/permissions
Authorization: Bearer <token>
```

## 使用示例

### 1. 初始化数据库

```bash
mysql -u root -p < backend/sql/init_rbac.sql
```

### 2. 创建新角色

```bash
curl -X POST http://localhost:8888/api/v1/roles \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "role_name": "编辑",
    "role_code": "editor",
    "description": "内容编辑角色"
  }'
```

### 3. 创建权限

```bash
curl -X POST http://localhost:8888/api/v1/permissions \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "perm_name": "发布文章",
    "perm_code": "article:publish",
    "perm_type": "button",
    "resource": "/api/v1/articles/publish",
    "action": "POST"
  }'
```

### 4. 分配权限给角色

```bash
curl -X POST http://localhost:8888/api/v1/roles/5/permissions \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "permission_ids": [10, 11, 12]
  }'
```

### 5. 分配角色给用户

```bash
curl -X POST http://localhost:8888/api/v1/users/2/roles \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "role_ids": [3, 5]
  }'
```

### 6. 查询用户权限

```bash
curl -X GET http://localhost:8888/api/v1/users/2/permissions \
  -H "Authorization: Bearer $TOKEN"
```

## 权限验证

### 后端中间件（规划中）

```go
// 权限验证中间件
func PermissionMiddleware(requiredPerm string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // 1. 从 context 获取用户 ID
            userId := r.Context().Value("userId").(int64)
            
            // 2. 查询用户权限
            permissions := GetUserPermissions(userId)
            
            // 3. 验证权限
            if !HasPermission(permissions, requiredPerm) {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}

// 使用示例
@handler CreateUser
post /api/v1/users (CreateUserReq) returns (CreateUserResp) {
    permission: "user:create"
}
```

### 前端权限控制（规划中）

```vue
<template>
  <div>
    <!-- 按钮级权限控制 -->
    <el-button v-permission="'user:create'">创建用户</el-button>
    
    <!-- 菜单级权限控制 -->
    <el-menu-item v-if="hasPermission('user:manage')" index="/users">
      用户管理
    </el-menu-item>
  </div>
</template>

<script setup>
import { hasPermission } from '@/utils/permission'

// 检查权限
if (hasPermission('user:delete')) {
  // 执行删除操作
}
</script>
```

## 默认权限配置

### 超级管理员 (super_admin)
- ✅ 所有权限

### 管理员 (admin)
- ✅ 用户管理（查看、创建、编辑、删除）
- ✅ 角色管理（查看、创建、编辑、删除、分配权限）
- ❌ 权限管理

### 普通用户 (user)
- ✅ 查看用户
- ✅ 查看角色

### 访客 (guest)
- ✅ 查看用户

## 最佳实践

### 1. 权限代码命名规范

```
资源：操作
例如:
- user:view      (查看用户)
- user:create    (创建用户)
- article:edit   (编辑文章)
- order:delete   (删除订单)
```

### 2. 角色设计原则

- **最小权限原则**: 角色只包含完成工作所需的最小权限
- **职责分离**: 不同角色之间权限不重叠
- **可扩展性**: 预留新增角色的空间

### 3. 安全建议

- 定期审计权限分配
- 敏感操作记录日志
- 限制超级管理员数量
- 定期轮换密钥

## 常见问题

### Q1: 如何添加新权限？

1. 在权限表中创建新权限记录
2. 将权限分配给需要的角色
3. 后端添加对应的权限验证
4. 前端添加对应的 UI 控制

### Q2: 用户有多个角色，权限如何计算？

用户的权限是所有角色权限的**并集**。

### Q3: 如何禁用某个权限？

1. 将权限状态设置为 0（禁用）
2. 或从所有角色中移除该权限

### Q4: 角色被删除后，关联的用户如何处理？

软删除角色，保留关联关系。或先转移用户到其他角色再删除。

---

**最后更新**: 2026-03-12
