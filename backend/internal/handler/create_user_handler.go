package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user-management-system/internal/logic"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
