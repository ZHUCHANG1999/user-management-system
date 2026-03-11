package logic

import (
	"context"
	"errors"
	"time"
	"user-management-system/internal/model"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.RoleCreateReq) (resp *types.RoleCreateResp, err error) {
	// 检查角色代码是否已存在
	existingRole, err := l.svcCtx.RoleModel.FindByCode(req.RoleCode)
	if err == nil && existingRole != nil {
		return nil, errors.New("角色代码已存在")
	}

	role := &model.Role{
		RoleName:    req.RoleName,
		RoleCode:    req.RoleCode,
		Description: req.Description,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := l.svcCtx.RoleModel.Create(role); err != nil {
		return nil, err
	}

	return &types.RoleCreateResp{
		RoleId:  role.RoleId,
		Message: "角色创建成功",
	}, nil
}

type GetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleLogic) GetRole(req *types.RoleGetReq) (resp *types.RoleGetResp, err error) {
	role, err := l.svcCtx.RoleModel.FindByID(req.RoleId)
	if err != nil {
		return nil, err
	}

	permissions := make([]types.PermissionInfo, 0, len(role.Permissions))
	for _, perm := range role.Permissions {
		permissions = append(permissions, types.PermissionInfo{
			PermissionId: perm.PermissionId,
			PermName:     perm.PermName,
			PermCode:     perm.PermCode,
			PermType:     perm.PermType,
			Resource:     perm.Resource,
			Action:       perm.Action,
			Description:  perm.Description,
			Status:       perm.Status,
		})
	}

	return &types.RoleGetResp{
		RoleId:      role.RoleId,
		RoleName:    role.RoleName,
		RoleCode:    role.RoleCode,
		Description: role.Description,
		Status:      role.Status,
		Permissions: permissions,
		CreatedAt:   role.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.RoleUpdateReq) (resp *types.RoleUpdateResp, err error) {
	role, err := l.svcCtx.RoleModel.FindByID(req.RoleId)
	if err != nil {
		return nil, err
	}

	role.RoleName = req.RoleName
	role.Description = req.Description
	if req.Status != 0 {
		role.Status = req.Status
	}

	if err := l.svcCtx.RoleModel.Update(role); err != nil {
		return nil, err
	}

	return &types.RoleUpdateResp{
		Message: "角色更新成功",
	}, nil
}

type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleLogic) DeleteRole(req *types.RoleDeleteReq) (resp *types.RoleDeleteResp, err error) {
	if err := l.svcCtx.RoleModel.Delete(req.RoleId); err != nil {
		return nil, err
	}

	return &types.RoleDeleteResp{
		Message: "角色删除成功",
	}, nil
}

type ListRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListRolesLogic {
	return &ListRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListRolesLogic) ListRoles(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}

	roles, total, err := l.svcCtx.RoleModel.FindPage(page, pageSize, req.RoleName)
	if err != nil {
		return nil, err
	}

	roleList := make([]types.RoleInfo, 0, len(roles))
	for _, role := range roles {
		roleList = append(roleList, types.RoleInfo{
			RoleId:      role.RoleId,
			RoleName:    role.RoleName,
			RoleCode:    role.RoleCode,
			Description: role.Description,
			Status:      role.Status,
			CreatedAt:   role.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &types.RoleListResp{
		Total: int(total),
		Roles: roleList,
	}, nil
}
