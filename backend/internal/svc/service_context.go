package svc

import (
	"user-management-system/internal/config"
	"user-management-system/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	DB             *gorm.DB
	UserModel      *model.UserDBModel
	RoleModel      *model.RoleDBModel
	PermissionModel *model.PermissionDBModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(c.Database.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 自动迁移表结构
	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Permission{})

	return &ServiceContext{
		Config:          c,
		DB:              db,
		UserModel:       model.NewUserModel(db),
		RoleModel:       model.NewRoleModel(db),
		PermissionModel: model.NewPermissionModel(db),
	}
}
