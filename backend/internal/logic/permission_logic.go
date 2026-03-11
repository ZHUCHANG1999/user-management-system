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

type CreatePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermissionLogic {
	return &CreatePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePermissionLogic) CreatePermission(req *types.PermissionCreateReq) (resp *types.PermissionCreateResp, err error) {
	// 检查权限代码是否已存在
	existingPerm, err := l.svcCtx.PermissionModel.FindByCode(req.PermCode)
	if err == nil && existingPerm != nil {
		return nil, errors.New("权限代码已存在")
	}

	perm := &model.Permission{
		PermName:    req.PermName,
		PermCode:    req.PermCode,
		PermType:    req.PermType,
		Resource:    req.Resource,
		Action:      req.Action,
		Description: req.Description,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := l.svcCtx.PermissionModel.Create(perm); err != nil {
		return nil, err
	}

	return &types.PermissionCreateResp{
		PermissionId: perm.PermissionId,
		Message:      "权限创建成功",
	}, nil
}

type GetPermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionLogic {
	return &GetPermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPermissionLogic) GetPermission(req *types.PermissionGetReq) (resp *types.PermissionGetResp, err error) {
	perm, err := l.svcCtx.PermissionModel.FindByID(req.PermissionId)
	if err != nil {
		return nil, err
	}

	return &types.PermissionGetResp{
		PermissionId: perm.PermissionId,
		PermName:     perm.PermName,
		PermCode:     perm.PermCode,
		PermType:     perm.PermType,
		Resource:     perm.Resource,
		Action:       perm.Action,
		Description:  perm.Description,
		Status:       perm.Status,
	}, nil
}

type UpdatePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePermissionLogic {
	return &UpdatePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePermissionLogic) UpdatePermission(req *types.PermissionUpdateReq) (resp *types.PermissionUpdateResp, err error) {
	perm, err := l.svcCtx.PermissionModel.FindByID(req.PermissionId)
	if err != nil {
		return nil, err
	}

	perm.PermName = req.PermName
	perm.Description = req.Description
	if req.Status != 0 {
		perm.Status = req.Status
	}

	if err := l.svcCtx.PermissionModel.Update(perm); err != nil {
		return nil, err
	}

	return &types.PermissionUpdateResp{
		Message: "权限更新成功",
	}, nil
}

type DeletePermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePermissionLogic {
	return &DeletePermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePermissionLogic) DeletePermission(req *types.PermissionDeleteReq) (resp *types.PermissionDeleteResp, err error) {
	if err := l.svcCtx.PermissionModel.Delete(req.PermissionId); err != nil {
		return nil, err
	}

	return &types.PermissionDeleteResp{
		Message: "权限删除成功",
	}, nil
}

type ListPermissionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPermissionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPermissionsLogic {
	return &ListPermissionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPermissionsLogic) ListPermissions(req *types.PermissionListReq) (resp *types.PermissionListResp, err error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}

	permissions, total, err := l.svcCtx.PermissionModel.FindPage(page, pageSize, req.PermType)
	if err != nil {
		return nil, err
	}

	permList := make([]types.PermissionInfo, 0, len(permissions))
	for _, perm := range permissions {
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

	return &types.PermissionListResp{
		Total:       int(total),
		Permissions: permList,
	}, nil
}
