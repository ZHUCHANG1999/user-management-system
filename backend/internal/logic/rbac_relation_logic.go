package logic

import (
	"context"
	"user-management-system/internal/svc"
	"user-management-system/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignPermissionsLogic {
	return &AssignPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignPermissionsLogic) AssignPermissions(req *types.AssignPermissionsReq) (resp *types.AssignPermissionsResp, err error) {
	if err := l.svcCtx.RoleModel.AssignPermissions(req.RoleId, req.PermissionIds); err != nil {
		return nil, err
	}

	return &types.AssignPermissionsResp{
		Message: "权限分配成功",
	}, nil
}

type GetRolePermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRolePermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolePermissionsLogic {
	return &GetRolePermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRolePermissionsLogic) GetRolePermissions(req *types.GetRolePermissionsReq) (resp *types.GetRolePermissionsResp, err error) {
	permissions, err := l.svcCtx.RoleModel.GetPermissions(req.RoleId)
	if err != nil {
		return nil, err
	}

	permissionIds := make([]int64, 0, len(permissions))
	permList := make([]types.PermissionInfo, 0, len(permissions))
	
	for _, perm := range permissions {
		permissionIds = append(permissionIds, perm.PermissionId)
		permList = append(permList, types.PermissionInfo{
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

	return &types.GetRolePermissionsResp{
		RoleId:        req.RoleId,
		PermissionIds: permissionIds,
		Permissions:   permList,
	}, nil
}

type AssignRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRolesLogic {
	return &AssignRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignRolesLogic) AssignRoles(req *types.AssignRolesReq) (resp *types.AssignRolesResp, err error) {
	// TODO: 实现用户 - 角色关联
	return &types.AssignRolesResp{
		Message: "角色分配成功",
	}, nil
}

type GetUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRolesLogic) GetUserRoles(req *types.GetUserRolesReq) (resp *types.GetUserRolesResp, err error) {
	// TODO: 实现用户角色查询
	return &types.GetUserRolesResp{
		UserId:  req.UserId,
		RoleIds: []int64{},
		Roles:   []types.RoleInfo{},
	}, nil
}

type GetUserPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPermissionsLogic {
	return &GetUserPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPermissionsLogic) GetUserPermissions(req *types.GetUserPermissionsReq) (resp *types.GetUserPermissionsResp, err error) {
	// TODO: 实现用户权限查询（通过角色关联）
	return &types.GetUserPermissionsResp{
		UserId:      req.UserId,
		Permissions: []types.PermissionInfo{},
	}, nil
}
