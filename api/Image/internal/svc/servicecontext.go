package svc

import (
	"coderhub/api/Image/internal/config"
	"coderhub/rpc/Image/image"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	ImageService image.Image
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ImageService: image.NewImage(zrpc.MustNewClient(c.ImageService)),
	}
}
