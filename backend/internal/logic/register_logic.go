package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"user-management-system/internal/svc"
	"user-management-system/types"
)

type RegisterLogic struct {
	ctx    context.Context
	serverCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, serverCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		serverCtx: serverCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.RegisterResponse, error) {
	// 参数校验
	if req.Username == "" || req.Password == "" || req.Email == "" {
		return nil, errors.New("用户名、密码和邮箱不能为空")
	}

	// 检查用户名是否已存在
	var count int
	err := l.serverCtx.SqlConn.QueryRowCtx(l.ctx, &count, 
		"SELECT COUNT(*) FROM `user` WHERE username = ? AND deleted_at IS NULL", req.Username)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	err = l.serverCtx.SqlConn.QueryRowCtx(l.ctx, &count,
		"SELECT COUNT(*) FROM `user` WHERE email = ? AND deleted_at IS NULL", req.Email)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("邮箱已被注册")
	}

	// 密码加密
	passwordHash := md5.Sum([]byte(req.Password))
	passwordHashStr := hex.EncodeToString(passwordHash[:])

	// 插入用户
	now := time.Now()
	result, err := l.serverCtx.SqlConn.ExecCtx(l.ctx,
		"INSERT INTO `user` (username, password_hash, email, nickname, role, status, created_at, updated_at) VALUES (?, ?, ?, ?, 'user', 1, ?, ?)",
		req.Username, passwordHashStr, req.Email, req.Nickname, now, now)
	if err != nil {
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		UserId:  userId,
		Message: "注册成功",
	}, nil
}
