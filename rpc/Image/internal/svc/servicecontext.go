package svc

import (
	repository "coderhub/repository/Image"
	"coderhub/rpc/Image/internal/config"
	"coderhub/shared/SQL"
)

type ServiceContext struct {
	Config          config.Config
	ImageRepository repository.ImageRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ImageRepository: repository.NewImageRepository(SQL.NewGorm()),
	}
}
