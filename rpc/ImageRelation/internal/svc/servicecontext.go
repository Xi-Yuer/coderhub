package svc

import (
	"coderhub/repository"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/internal/config"
	"coderhub/shared/storage"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                  config.Config
	ImageRepository         repository.ImageRepository
	ImageRelationRepository repository.ImageRelationRepository
	ImageService            imageservice.ImageService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                  c,
		ImageRepository:         repository.NewImageRepository(storage.NewGorm()),
		ImageRelationRepository: repository.NewImageRelationRepository(storage.NewGorm()),
		ImageService:            imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
	}
}
