package logic

import (
	"context"
	"errors"
	"time"

	"user-management-system/internal/svc"
	"user-management-system/types"

	"github.com/golang-jwt/jwt/v4"
)

type DeleteUserLogic struct {
	ctx    context.Context
	serverCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, serverCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		serverCtx: serverCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserRequest) (*types.CommonResponse, error) {
	// 从 JWT Token 中获取当前用户信息和角色
	tokenStr := l.ctx.Value("token")
	if tokenStr == nil {
		return nil, errors.New("未授权")
	}

	token, err := jwt.ParseWithClaims(tokenStr.(string), &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(l.serverCtx.Config.Auth.AccessSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("Token 无效")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("Token 解析失败")
	}

	// 只有管理员可以删除用户
	if claims.Role != "admin" {
		return nil, errors.New("权限不足，只有管理员可以删除用户")
	}

	// 不能删除自己
	if claims.UserId == req.UserId {
		return nil, errors.New("不能删除自己的账号")
	}

	// 检查要删除的用户是否存在且不是管理员
	var role string
	err = l.serverCtx.SqlConn.QueryRowCtx(l.ctx, &role,
		"SELECT role FROM `user` WHERE user_id = ? AND deleted_at IS NULL", req.UserId)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	if role == "admin" {
		return nil, errors.New("不能删除管理员账号")
	}

	// 软删除
	_, err = l.serverCtx.SqlConn.ExecCtx(l.ctx,
		"UPDATE `user` SET deleted_at = ? WHERE user_id = ?",
		time.Now(), req.UserId)
	if err != nil {
		return nil, err
	}

	return &types.CommonResponse{
		Message: "删除成功",
	}, nil
}
