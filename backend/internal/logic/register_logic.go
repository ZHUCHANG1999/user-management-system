package logic

import (
	"context"
	"errors"
	"time"
	"user-management-system/internal/model"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 检查用户名是否已存在
	existingUser, err := l.svcCtx.UserModel.FindByUsername(req.Username)
	if err == nil && existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户对象
	user := &model.User{
		Username:  req.Username,
		Password:  string(hashedPassword),
		Email:     req.Email,
		Nickname:  req.Nickname,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 保存到数据库
	if err := l.svcCtx.UserModel.Create(user); err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		UserId:  user.UserId,
		Message: "注册成功",
	}, nil
}
