package logic

import (
	"context"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	// 查询用户
	user, err := l.svcCtx.UserModel.FindByID(req.UserId)
	if err != nil {
		return nil, err
	}

	// 更新用户信息
	user.Email = req.Email
	user.Nickname = req.Nickname
	if req.Status != nil {
		user.Status = *req.Status
	}

	// 保存到数据库
	if err := l.svcCtx.UserModel.Update(user); err != nil {
		return nil, err
	}

	return &types.UserUpdateResp{
		Message: "用户更新成功",
	}, nil
}
