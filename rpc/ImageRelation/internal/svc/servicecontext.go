package svc

import (
	Imagerepository "coderhub/repository/Image"
	ImagerelationRepository "coderhub/repository/ImageRelation"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/ImageRelation/internal/config"
	"coderhub/shared/SQL"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                  config.Config
	ImageRepository         Imagerepository.ImageRepository
	ImageRelationRepository ImagerelationRepository.ImageRelationRepository
	ImageService            imagerelationservice.ImageRelationService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                  c,
		ImageRepository:         Imagerepository.NewImageRepository(SQL.NewGorm()),
		ImageRelationRepository: ImagerelationRepository.NewImageRelationRepository(SQL.NewGorm()),
		ImageService:            imagerelationservice.NewImageRelationService(zrpc.MustNewClient(c.ImageService)),
	}
}
