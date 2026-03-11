package logic

import (
	"context"
	"time"

	"user-management-system/internal/svc"
	"user-management-system/types"
)

type UserListLogic struct {
	ctx    context.Context
	serverCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, serverCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		serverCtx: serverCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserListRequest) (*types.UserListResponse, error) {
	// 计算偏移量
	offset := (req.Page - 1) * req.PageSize

	// 查询总数
	var total int64
	err := l.serverCtx.SqlConn.QueryRowCtx(l.ctx, &total,
		"SELECT COUNT(*) FROM `user` WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}

	// 查询用户列表
	rows, err := l.serverCtx.SqlConn.QueryCtx(l.ctx,
		"SELECT user_id, username, email, nickname, avatar, role, created_at FROM `user` WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT ? OFFSET ?",
		req.PageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []types.UserInfo
	for rows.Next() {
		var user struct {
			UserId    int64
			Username  string
			Email     string
			Nickname  string
			Avatar    string
			Role      string
			CreatedAt time.Time
		}
		err := rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Nickname, &user.Avatar, &user.Role, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, types.UserInfo{
			UserId:    user.UserId,
			Username:  user.Username,
			Email:     user.Email,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.Unix(),
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &types.UserListResponse{
		Total: total,
		Users: users,
	}, nil
}
