package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	UserId    int64     `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	Username  string    `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Email     string    `gorm:"size:100" json:"email"`
	Nickname  string    `gorm:"size:50" json:"nickname"`
	Status    int       `gorm:"default:1" json:"status"` // 1: 正常，0: 禁用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// UserDBModel 数据库操作模型
type UserDBModel struct {
	db *gorm.DB
}

// NewUserModel 创建用户模型实例
func NewUserModel(db *gorm.DB) *UserDBModel {
	return &UserDBModel{db: db}
}

// Create 创建用户
func (m *UserDBModel) Create(user *User) error {
	return m.db.Create(user).Error
}

// FindByID 根据 ID 查询用户
func (m *UserDBModel) FindByID(id int64) (*User, error) {
	var user User
	err := m.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername 根据用户名查询用户
func (m *UserDBModel) FindByUsername(username string) (*User, error) {
	var user User
	err := m.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update 更新用户
func (m *UserDBModel) Update(user *User) error {
	return m.db.Save(user).Error
}

// Delete 删除用户（软删除）
func (m *UserDBModel) Delete(id int64) error {
	return m.db.Delete(&User{}, id).Error
}

// FindPage 分页查询用户列表
func (m *UserDBModel) FindPage(page, pageSize int, username string) ([]User, int64, error) {
	var users []User
	var total int64

	query := m.db.Model(&User{})
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
