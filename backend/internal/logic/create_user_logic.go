package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
	"user-management-system/internal/model"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
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

	return &types.UserCreateResp{
		UserId:  user.UserId,
		Message: "用户创建成功",
	}, nil
}
