package svc

import (
	"coderhub/api/user/internal/config"
	"coderhub/rpc/user/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
	}
}
