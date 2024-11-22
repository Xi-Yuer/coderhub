package svc

import (
	repository "coderhub/repository/User"
	"coderhub/rpc/User/internal/config"
	"coderhub/shared/CacheDB"
	"coderhub/shared/SQL"
)

type ServiceContext struct {
	Config         config.Config
	UserRepository repository.UserRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := CacheDB.NewRedisDB(CacheDB.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:         c,
		UserRepository: repository.NewUserRepositoryImpl(SQL.NewGorm(), redisDB),
	}
}
