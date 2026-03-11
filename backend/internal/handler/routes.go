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
	server.AddRoutes(
		[]httpx.Route{
			{
				Method:   http.MethodPost,
				Path:     "/api/v1/users",
				Handler:  CreateUserHandler(ctx),
				Middleware: []httpx.Middleware{middleware.AuthMiddleware(ctx.Config.JWT.Secret)},
			},
			{
				Method:   http.MethodGet,
				Path:     "/api/v1/users/:user_id",
				Handler:  GetUserHandler(ctx),
				Middleware: []httpx.Middleware{middleware.AuthMiddleware(ctx.Config.JWT.Secret)},
			},
			{
				Method:   http.MethodPut,
				Path:     "/api/v1/users/:user_id",
				Handler:  UpdateUserHandler(ctx),
				Middleware: []httpx.Middleware{middleware.AuthMiddleware(ctx.Config.JWT.Secret)},
			},
			{
				Method:   http.MethodDelete,
				Path:     "/api/v1/users/:user_id",
				Handler:  DeleteUserHandler(ctx),
				Middleware: []httpx.Middleware{middleware.AuthMiddleware(ctx.Config.JWT.Secret)},
			},
			{
				Method:   http.MethodGet,
				Path:     "/api/v1/users",
				Handler:  ListUsersHandler(ctx),
				Middleware: []httpx.Middleware{middleware.AuthMiddleware(ctx.Config.JWT.Secret)},
			},
		},
	)
}
