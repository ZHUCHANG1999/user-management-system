package logic

import (
	"context"
	"errors"
	"time"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"
	"user-management-system/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 查询用户
	user, err := l.svcCtx.UserModel.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if user.Status != 1 {
		return nil, errors.New("用户已被禁用")
	}

	// 生成 Token
	token, err := utils.GenerateToken(
		user.UserId,
		user.Username,
		l.svcCtx.Config.JWT.Secret,
		l.svcCtx.Config.JWT.Expire,
	)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Token:  token,
		Expire: l.svcCtx.Config.JWT.Expire,
		User: types.UserInfo{
			UserId:    user.UserId,
			Username:  user.Username,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}
