package model

import (
	"time"

	"gorm.io/gorm"
)

// Role 角色模型
type Role struct {
	RoleId      int64     `gorm:"primaryKey;autoIncrement;column:role_id" json:"role_id"`
	RoleName    string    `gorm:"size:50;not null" json:"role_name"`
	RoleCode    string    `gorm:"uniqueIndex;size:50;not null" json:"role_code"`
	Description string    `gorm:"size:200" json:"description"`
	Status      int       `gorm:"default:1" json:"status"` // 1: 正常，0: 禁用
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	
	// 关联
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

// TableName 指定表名
func (Role) TableName() string {
	return "roles"
}

// Permission 权限模型
type Permission struct {
	PermissionId int64     `gorm:"primaryKey;autoIncrement;column:permission_id" json:"permission_id"`
	PermName     string    `gorm:"size:50;not null" json:"perm_name"`
	PermCode     string    `gorm:"uniqueIndex;size:100;not null" json:"perm_code"`
	PermType     string    `gorm:"size:20;not null" json:"perm_type"` // menu, button, api
	Resource     string    `gorm:"size:100" json:"resource"`
	Action       string    `gorm:"size:50" json:"action"`
	Description  string    `gorm:"size:200" json:"description"`
	Status       int       `gorm:"default:1" json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (Permission) TableName() string {
	return "permissions"
}

// RoleDBModel 角色数据库操作
type RoleDBModel struct {
	db *gorm.DB
}

// NewRoleModel 创建角色模型实例
func NewRoleModel(db *gorm.DB) *RoleDBModel {
	return &RoleDBModel{db: db}
}

// Create 创建角色
func (m *RoleDBModel) Create(role *Role) error {
	return m.db.Create(role).Error
}

// FindByID 根据 ID 查询角色
func (m *RoleDBModel) FindByID(id int64) (*Role, error) {
	var role Role
	err := m.db.Preload("Permissions").First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// FindByCode 根据角色代码查询角色
func (m *RoleDBModel) FindByCode(code string) (*Role, error) {
	var role Role
	err := m.db.Where("role_code = ?", code).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// Update 更新角色
func (m *RoleDBModel) Update(role *Role) error {
	return m.db.Save(role).Error
}

// Delete 删除角色（软删除）
func (m *RoleDBModel) Delete(id int64) error {
	return m.db.Delete(&Role{}, id).Error
}

// FindPage 分页查询角色列表
func (m *RoleDBModel) FindPage(page, pageSize int, roleName string) ([]Role, int64, error) {
	var roles []Role
	var total int64

	query := m.db.Model(&Role{})
	if roleName != "" {
		query = query.Where("role_name LIKE ?", "%"+roleName+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

// AssignPermissions 分配权限给角色
func (m *RoleDBModel) AssignPermissions(roleId int64, permissionIds []int64) error {
	var role Role
	if err := m.db.First(&role, roleId).Error; err != nil {
		return err
	}

	var permissions []Permission
	if err := m.db.Find(&permissions, permissionIds).Error; err != nil {
		return err
	}

	return m.db.Model(&role).Association("Permissions").Replace(permissions)
}

// GetPermissions 获取角色的权限
func (m *RoleDBModel) GetPermissions(roleId int64) ([]Permission, error) {
	var role Role
	if err := m.db.First(&role, roleId).Error; err != nil {
		return nil, err
	}

	var permissions []Permission
	err := m.db.Model(&role).Association("Permissions").Find(&permissions)
	return permissions, err
}

// PermissionDBModel 权限数据库操作
type PermissionDBModel struct {
	db *gorm.DB
}

// NewPermissionModel 创建权限模型实例
func NewPermissionModel(db *gorm.DB) *PermissionDBModel {
	return &PermissionDBModel{db: db}
}

// Create 创建权限
func (m *PermissionDBModel) Create(perm *Permission) error {
	return m.db.Create(perm).Error
}

// FindByID 根据 ID 查询权限
func (m *PermissionDBModel) FindByID(id int64) (*Permission, error) {
	var perm Permission
	err := m.db.First(&perm, id).Error
	if err != nil {
		return nil, err
	}
	return &perm, nil
}

// Update 更新权限
func (m *PermissionDBModel) Update(perm *Permission) error {
	return m.db.Save(perm).Error
}

// Delete 删除权限（软删除）
func (m *PermissionDBModel) Delete(id int64) error {
	return m.db.Delete(&Permission{}, id).Error
}

// FindPage 分页查询权限列表
func (m *PermissionDBModel) FindPage(page, pageSize int, permType string) ([]Permission, int64, error) {
	var permissions []Permission
	var total int64

	query := m.db.Model(&Permission{})
	if permType != "" {
		query = query.Where("perm_type = ?", permType)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&permissions).Error
	if err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}

// FindByCode 根据权限代码查询
func (m *PermissionDBModel) FindByCode(code string) (*Permission, error) {
	var perm Permission
	err := m.db.Where("perm_code = ?", code).First(&perm).Error
	if err != nil {
		return nil, err
	}
	return &perm, nil
}

// FindByCodes 根据权限代码批量查询
func (m *PermissionDBModel) FindByCodes(codes []string) ([]Permission, error) {
	var permissions []Permission
	err := m.db.Where("perm_code IN ?", codes).Find(&permissions).Error
	return permissions, err
}

// FindAll 查询所有权限
func (m *PermissionDBModel) FindAll() ([]Permission, error) {
	var permissions []Permission
	err := m.db.Order("perm_type, perm_code").Find(&permissions).Error
	return permissions, err
}
