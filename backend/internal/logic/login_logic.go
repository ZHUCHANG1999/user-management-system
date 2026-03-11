package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"

	"user-management-system/internal/svc"
	"user-management-system/types"

	"github.com/golang-jwt/jwt/v4"
)

type LoginLogic struct {
	ctx    context.Context
	serverCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, serverCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
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

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, error) {
	// 参数校验
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名和密码不能为空")
	}

	// 查询用户
	var user struct {
		UserId       int64
		Username     string
		PasswordHash string
		Role         string
		Status       int
	}

	err := l.serverCtx.SqlConn.QueryRowCtx(l.ctx, &user,
		"SELECT user_id, username, password_hash, role, status FROM `user` WHERE username = ? AND deleted_at IS NULL",
		req.Username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	// 检查用户状态
	if user.Status != 1 {
		return nil, errors.New("用户已被禁用")
	}

	// 验证密码
	passwordHash := md5.Sum([]byte(req.Password))
	passwordHashStr := hex.EncodeToString(passwordHash[:])
	if passwordHashStr != user.PasswordHash {
		return nil, errors.New("密码错误")
	}

	// 生成 JWT Token
	expireTime := time.Now().Add(time.Duration(l.serverCtx.Config.Auth.AccessExpire) * time.Second)
	claims := Claims{
		UserId:   user.UserId,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "user-management-system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(l.serverCtx.Config.Auth.AccessSecret))
	if err != nil {
		return nil, err
	}

	return &types.LoginResponse{
		UserId: user.UserId,
		Token:  tokenString,
		Expire: expireTime.Unix(),
	}, nil
}
