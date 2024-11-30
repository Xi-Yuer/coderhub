package svc

import (
	"coderhub/api/TechSphere/Comment/internal/config"
	"coderhub/rpc/TechSphere/Comment/commentservice"
	"coderhub/rpc/User/userservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	CommentService commentservice.CommentService
	UserService    userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CommentService: commentservice.NewCommentService(zrpc.MustNewClient(c.CommentService)),
	}
}
