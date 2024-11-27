package svc

import (
	"coderhub/api/TechSphere/Comment/internal/config"
	"coderhub/rpc/TechSphere/Comment/commentservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	CommentService commentservice.CommentService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CommentService: commentservice.NewCommentService(zrpc.MustNewClient(c.CommentService)),
	}
}
