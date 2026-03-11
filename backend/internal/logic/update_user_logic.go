package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"user-management-system/internal/svc"
	"user-management-system/types"

	"github.com/golang-jwt/jwt/v4"
)

type UpdateUserLogic struct {
	ctx    context.Context
	serverCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, serverCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		serverCtx: serverCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserRequest) (*types.CommonResponse, error) {
	// 从 JWT Token 中获取用户 ID
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

	// 构建更新字段
	var updateFields []string
	var updateValues []interface{}

	if req.Nickname != "" {
		updateFields = append(updateFields, "nickname = ?")
		updateValues = append(updateValues, req.Nickname)
	}

	if req.Email != "" {
		// 检查邮箱是否已被其他用户使用
		var count int
		err = l.serverCtx.SqlConn.QueryRowCtx(l.ctx, &count,
			"SELECT COUNT(*) FROM `user` WHERE email = ? AND user_id != ? AND deleted_at IS NULL",
			req.Email, claims.UserId)
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, errors.New("邮箱已被使用")
		}
		updateFields = append(updateFields, "email = ?")
		updateValues = append(updateValues, req.Email)
	}

	if req.Avatar != "" && strings.HasPrefix(req.Avatar, "http") {
		updateFields = append(updateFields, "avatar = ?")
		updateValues = append(updateValues, req.Avatar)
	}

	if len(updateFields) == 0 {
		return nil, errors.New("没有需要更新的字段")
	}

	// 添加更新时间和用户 ID
	updateFields = append(updateFields, "updated_at = ?")
	updateValues = append(updateValues, time.Now())
	updateValues = append(updateValues, claims.UserId)

	// 执行更新
	query := "UPDATE `user` SET " + strings.Join(updateFields, ", ") + " WHERE user_id = ? AND deleted_at IS NULL"
	_, err = l.serverCtx.SqlConn.ExecCtx(l.ctx, query, updateValues...)
	if err != nil {
		return nil, err
	}

	return &types.CommonResponse{
		Message: "更新成功",
	}, nil
}
