package logic

import (
	"context"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserGetReq) (resp *types.UserGetResp, err error) {
	// 从数据库查询用户
	user, err := l.svcCtx.UserModel.FindByID(req.UserId)
	if err != nil {
		return nil, err
	}

	return &types.UserGetResp{
		UserId:    user.UserId,
		Username:  user.Username,
		Email:     user.Email,
		Nickname:  user.Nickname,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
