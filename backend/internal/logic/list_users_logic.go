package logic

import (
	"context"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListUsersLogic) ListUsers(req *types.UserListReq) (resp *types.UserListResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}

	// 从数据库查询用户列表
	users, total, err := l.svcCtx.UserModel.FindPage(page, pageSize, req.Username)
	if err != nil {
		return nil, err
	}

	// 转换数据格式
	userList := make([]types.UserInfo, 0, len(users))
	for _, user := range users {
		userList = append(userList, types.UserInfo{
			UserId:    user.UserId,
			Username:  user.Username,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Status:    user.Status,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.UserListResp{
		Total: int(total),
		Users: userList,
	}, nil
}
