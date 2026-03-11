package svc

import "user-management-system/internal/model"

type ServiceContext struct {
	// 后续添加数据库连接等
}

func NewServiceContext() *ServiceContext {
	return &ServiceContext{}
}
