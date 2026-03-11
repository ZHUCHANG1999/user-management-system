package model

import "time"

// User 用户模型
type User struct {
	UserId    int64     `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username  string    `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"-"`
	Email     string    `gorm:"size:100" json:"email"`
	Nickname  string    `gorm:"size:50" json:"nickname"`
	Status    int       `gorm:"default:1" json:"status"` // 1: 正常，0: 禁用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
