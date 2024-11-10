package svc

import (
	repository "coderhub/repository/user"
	"coderhub/rpc/user/internal/config"
	"coderhub/shared/cacheDB"
	"coderhub/shared/sqlDB"
)

type ServiceContext struct {
	Config         config.Config
	UserRepository repository.UserRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRepository: repository.NewUserRepositoryImpl(sqlDB.NewGorm(), cacheDB.NewRedisDB()),
	}
}
