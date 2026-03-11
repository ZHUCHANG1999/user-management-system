package logic

import (
	"context"
	"errors"
	"time"

	"user-management-system/internal/svc"
	"user-management-system/types"

	"github.com/golang-jwt/jwt/v4"
)

type GetUserLogic struct {
	ctx    context.Context
	serverCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, serverCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		serverCtx: serverCtx,
	}
}

type Claims struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (*types.GetUserResponse, error) {
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

	// 查询用户信息
	var user struct {
		UserId    int64
		Username  string
		Email     string
		Nickname  string
		Avatar    string
		Role      string
		CreatedAt time.Time
	}

	err = l.serverCtx.SqlConn.QueryRowCtx(l.ctx, &user,
		"SELECT user_id, username, email, nickname, avatar, role, created_at FROM `user` WHERE user_id = ? AND deleted_at IS NULL",
		claims.UserId)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	return &types.GetUserResponse{
		User: types.UserInfo{
			UserId:    user.UserId,
			Username:  user.Username,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Unix(),
		},
	}, nil
}
