package handler

import (
	"net/http"

	"user-management-system/internal/svc"
)

func RegisterHandlers(server httpx.Server, ctx *svc.ServiceContext) {
	server.AddRoutes(
		[]httpx.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/users",
				Handler: CreateUserHandler(ctx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/users/:user_id",
				Handler: GetUserHandler(ctx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/v1/users/:user_id",
				Handler: UpdateUserHandler(ctx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/v1/users/:user_id",
				Handler: DeleteUserHandler(ctx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/users",
				Handler: ListUsersHandler(ctx),
			},
		},
	)
}
