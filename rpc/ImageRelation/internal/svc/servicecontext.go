package svc

import (
	"coderhub/repository"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/internal/config"
	"coderhub/shared/SQL"

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
		ImageRepository:         repository.NewImageRepository(SQL.NewGorm()),
		ImageRelationRepository: repository.NewImageRelationRepository(SQL.NewGorm()),
		ImageService:            imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
	}
}
