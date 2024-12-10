package svc

import (
	"coderhub/repository"
	"coderhub/rpc/User/userservice"
	"coderhub/rpc/UserFollow/internal/config"
	"coderhub/shared/storage"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	UserService          userservice.UserService
	UserFollowRepository repository.UserFollowRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := storage.NewRedisDB(storage.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:               c,
		UserService:          userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
		UserFollowRepository: repository.NewUserFollowRepositoryImpl(storage.NewGorm(), redisDB),
	}
}
