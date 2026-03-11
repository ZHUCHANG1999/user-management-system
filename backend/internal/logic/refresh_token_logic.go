package logic

import (
	"context"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"
	"user-management-system/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.RefreshTokenReq) (resp *types.RefreshTokenResp, err error) {
	// 刷新 Token
	newToken, err := utils.RefreshToken(
		req.Token,
		l.svcCtx.Config.JWT.Secret,
		l.svcCtx.Config.JWT.Expire,
	)
	if err != nil {
		return nil, err
	}

	return &types.RefreshTokenResp{
		Token:  newToken,
		Expire: l.svcCtx.Config.JWT.Expire,
	}, nil
}
