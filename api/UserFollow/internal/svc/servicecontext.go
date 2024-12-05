package svc

import (
	"coderhub/api/UserFollow/internal/config"
	"coderhub/rpc/UserFollow/userfollowservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	UserFollowService userfollowservice.UserFollowService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UserFollowService: userfollowservice.NewUserFollowService(zrpc.MustNewClient(c.UserFollowService)),
	}
}
