package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user-management-system/internal/utils"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 获取 Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				httpx.Error(w, &httpx.Error{
					Err:  "missing authorization header",
					Code: http.StatusUnauthorized,
				})
				return
			}

			// 提取 Token
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				httpx.Error(w, &httpx.Error{
					Err:  "invalid authorization format",
					Code: http.StatusUnauthorized,
				})
				return
			}

			tokenString := parts[1]

			// 验证 Token
			claims, err := utils.ParseToken(tokenString, secret)
			if err != nil {
				httpx.Error(w, &httpx.Error{
					Err:  err.Error(),
					Code: http.StatusUnauthorized,
				})
				return
			}

			// 将用户信息存入 context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userId", claims.UserId)
			ctx = context.WithValue(ctx, "username", claims.Username)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
