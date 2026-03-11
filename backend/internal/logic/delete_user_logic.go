package logic

import (
	"context"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.UserDeleteReq) (resp *types.UserDeleteResp, err error) {
	// 软删除用户
	if err := l.svcCtx.UserModel.Delete(req.UserId); err != nil {
		return nil, err
	}

	return &types.UserDeleteResp{
		Message: "用户删除成功",
	}, nil
}
