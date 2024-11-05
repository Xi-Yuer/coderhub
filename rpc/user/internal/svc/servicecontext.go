package svc

import (
	"coderhub/rpc/user/internal/config"
	"coderhub/shared/sqlDB"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	SqlDB  *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SqlDB:  sqlDB.NewGorm(),
	}
}
