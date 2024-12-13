package svc

import (
	"coderhub/repository"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/User/internal/config"
	"coderhub/shared/messaging"
	"coderhub/shared/storage"
	"coderhub/shared/utils"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	Validator            *utils.Validator
	ImageRelationService imagerelationservice.ImageRelationService
	ImageService         imageservice.ImageService
	UserRepository       repository.UserRepository
	RedisDB              storage.RedisDB
	GoMail               messaging.GoMailImpl
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := storage.NewRedisDB(storage.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:               c,
		Validator:            utils.NewValidator(),
		ImageRelationService: imagerelationservice.NewImageRelationService(zrpc.MustNewClient(c.ImageRelationService)),
		ImageService:         imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
		RedisDB:              redisDB,
		GoMail:               messaging.NewGoMail(),
		UserRepository:       repository.NewUserRepositoryImpl(storage.NewGorm(), redisDB),
	}
}
