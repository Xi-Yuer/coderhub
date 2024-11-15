package svc

import (
	"coderhub/api/TechSphere/Articles/internal/config"
	"coderhub/rpc/TechSphere/Articles/articleservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	ArticleService articleservice.ArticleService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		ArticleService: articleservice.NewArticleService(zrpc.MustNewClient(c.ArticleService)),
	}
}
