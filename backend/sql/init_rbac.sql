-- RBAC 权限管理初始化脚本

USE `user_management`;

-- 角色表
CREATE TABLE IF NOT EXISTS `roles` (
  `role_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '角色 ID',
  `role_name` varchar(50) NOT NULL COMMENT '角色名称',
  `role_code` varchar(50) NOT NULL COMMENT '角色代码',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：1-正常，0-禁用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`),
  UNIQUE KEY `uk_role_code` (`role_code`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 权限表
CREATE TABLE IF NOT EXISTS `permissions` (
  `permission_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '权限 ID',
  `perm_name` varchar(50) NOT NULL COMMENT '权限名称',
  `perm_code` varchar(100) NOT NULL COMMENT '权限代码',
  `perm_type` varchar(20) NOT NULL COMMENT '权限类型：menu-菜单，button-按钮，api-接口',
  `resource` varchar(100) DEFAULT NULL COMMENT '资源路径',
  `action` varchar(50) DEFAULT NULL COMMENT '操作类型',
  `description` varchar(200) DEFAULT NULL COMMENT '描述',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：1-正常，0-禁用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`permission_id`),
  UNIQUE KEY `uk_perm_code` (`perm_code`),
  KEY `idx_type` (`perm_type`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- 角色 - 权限关联表
CREATE TABLE IF NOT EXISTS `role_permissions` (
  `role_id` bigint(20) NOT NULL COMMENT '角色 ID',
  `permission_id` bigint(20) NOT NULL COMMENT '权限 ID',
  PRIMARY KEY (`role_id`, `permission_id`),
  KEY `idx_permission_id` (`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- 用户 - 角色关联表
CREATE TABLE IF NOT EXISTS `user_roles` (
  `user_id` bigint(20) NOT NULL COMMENT '用户 ID',
  `role_id` bigint(20) NOT NULL COMMENT '角色 ID',
  PRIMARY KEY (`user_id`, `role_id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- 插入默认角色
INSERT INTO `roles` (`role_name`, `role_code`, `description`, `status`) VALUES
('超级管理员', 'super_admin', '拥有所有权限', 1),
('管理员', 'admin', '拥有管理权限', 1),
('普通用户', 'user', '基础用户权限', 1),
('访客', 'guest', '只读权限', 1);

-- 插入默认权限
INSERT INTO `permissions` (`perm_name`, `perm_code`, `perm_type`, `resource`, `action`, `description`) VALUES
-- 用户管理
('用户管理', 'user:manage', 'menu', '/users', '', '用户管理菜单'),
('查看用户', 'user:view', 'api', '/api/v1/users', 'GET', '查看用户列表和详情'),
('创建用户', 'user:create', 'api', '/api/v1/users', 'POST', '创建新用户'),
('编辑用户', 'user:edit', 'api', '/api/v1/users/:id', 'PUT', '编辑用户信息'),
('删除用户', 'user:delete', 'api', '/api/v1/users/:id', 'DELETE', '删除用户'),

-- 角色管理
('角色管理', 'role:manage', 'menu', '/roles', '', '角色管理菜单'),
('查看角色', 'role:view', 'api', '/api/v1/roles', 'GET', '查看角色列表'),
('创建角色', 'role:create', 'api', '/api/v1/roles', 'POST', '创建新角色'),
('编辑角色', 'role:edit', 'api', '/api/v1/roles/:id', 'PUT', '编辑角色'),
('删除角色', 'role:delete', 'api', '/api/v1/roles/:id', 'DELETE', '删除角色'),
('分配权限', 'role:assign', 'api', '/api/v1/roles/:id/permissions', 'POST', '为角色分配权限'),

-- 权限管理
('权限管理', 'permission:manage', 'menu', '/permissions', '', '权限管理菜单'),
('查看权限', 'permission:view', 'api', '/api/v1/permissions', 'GET', '查看权限列表'),
('创建权限', 'permission:create', 'api', '/api/v1/permissions', 'POST', '创建新权限'),
('编辑权限', 'permission:edit', 'api', '/api/v1/permissions/:id', 'PUT', '编辑权限'),
('删除权限', 'permission:delete', 'api', '/api/v1/permissions/:id', 'DELETE', '删除权限');

-- 为超级管理员分配所有权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`)
SELECT 1, permission_id FROM permissions;

-- 为管理员分配大部分权限（除权限管理）
INSERT INTO `role_permissions` (`role_id`, `permission_id`)
SELECT 2, permission_id FROM permissions WHERE perm_code NOT LIKE 'permission:%';

-- 为普通用户分配基础权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`)
SELECT 3, permission_id FROM permissions WHERE perm_code IN ('user:view', 'role:view');

-- 为访客分配只读权限
INSERT INTO `role_permissions` (`role_id`, `permission_id`)
SELECT 4, permission_id FROM permissions WHERE perm_code IN ('user:view');

-- 将 admin 用户设置为超级管理员
INSERT INTO `user_roles` (`user_id`, `role_id`)
SELECT user_id, 1 FROM users WHERE username = 'admin';

-- 查看表结构
DESC `roles`;
DESC `permissions`;
DESC `role_permissions`;
DESC `user_roles`;

-- 查看测试数据
SELECT * FROM `roles`;
SELECT * FROM `permissions`;
SELECT * FROM `role_permissions`;
SELECT * FROM `user_roles`;
