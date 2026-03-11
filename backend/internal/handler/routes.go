package handler

import (
	"net/http"

	"user-management-system/internal/middleware"
	"user-management-system/internal/svc"
)

func RegisterHandlers(server httpx.Server, ctx *svc.ServiceContext) {
	// 公开接口（无需认证）
	server.AddRoutes(
		[]httpx.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/auth/login",
				Handler: LoginHandler(ctx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/auth/register",
				Handler: RegisterHandler(ctx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/auth/refresh",
				Handler: RefreshTokenHandler(ctx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/auth/logout",
				Handler: LogoutHandler(ctx),
			},
		},
	)

	// 需要认证的接口
	authMiddleware := middleware.AuthMiddleware(ctx.Config.JWT.Secret)
	
	// 用户管理
	server.AddRoutes(
		[]httpx.Route{
			{
				Method:     http.MethodPost,
				Path:       "/api/v1/users",
				Handler:    CreateUserHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/users/:user_id",
				Handler:    GetUserHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodPut,
				Path:       "/api/v1/users/:user_id",
				Handler:    UpdateUserHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodDelete,
				Path:       "/api/v1/users/:user_id",
				Handler:    DeleteUserHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/users",
				Handler:    ListUsersHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
		},
	)

	// RBAC 角色权限管理
	server.AddRoutes(
		[]httpx.Route{
			// 角色管理
			{
				Method:     http.MethodPost,
				Path:       "/api/v1/roles",
				Handler:    CreateRoleHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/roles/:role_id",
				Handler:    GetRoleHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodPut,
				Path:       "/api/v1/roles/:role_id",
				Handler:    UpdateRoleHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodDelete,
				Path:       "/api/v1/roles/:role_id",
				Handler:    DeleteRoleHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/roles",
				Handler:    ListRolesHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			// 权限管理
			{
				Method:     http.MethodPost,
				Path:       "/api/v1/permissions",
				Handler:    CreatePermissionHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/permissions/:permission_id",
				Handler:    GetPermissionHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodPut,
				Path:       "/api/v1/permissions/:permission_id",
				Handler:    UpdatePermissionHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodDelete,
				Path:       "/api/v1/permissions/:permission_id",
				Handler:    DeletePermissionHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/permissions",
				Handler:    ListPermissionsHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			// 角色 - 权限关联
			{
				Method:     http.MethodPost,
				Path:       "/api/v1/roles/:role_id/permissions",
				Handler:    AssignPermissionsHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/roles/:role_id/permissions",
				Handler:    GetRolePermissionsHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			// 用户 - 角色关联
			{
				Method:     http.MethodPost,
				Path:       "/api/v1/users/:user_id/roles",
				Handler:    AssignRolesHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/users/:user_id/roles",
				Handler:    GetUserRolesHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
			{
				Method:     http.MethodGet,
				Path:       "/api/v1/users/:user_id/permissions",
				Handler:    GetUserPermissionsHandler(ctx),
				Middleware: []httpx.Middleware{authMiddleware},
			},
		},
	)
}
