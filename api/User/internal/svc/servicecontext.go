package svc

import (
	"coderhub/api/User/internal/config"
	"coderhub/rpc/User/userservice"

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
