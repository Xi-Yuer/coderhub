package svc

import (
	"coderhub/api/Image/internal/config"
	"coderhub/rpc/Image/imageservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	ImageService imageservice.ImageService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ImageService: imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
	}
}
