package svc

import (
	"coderhub/repository"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/User/internal/config"
	"coderhub/shared/CacheDB"
	"coderhub/shared/GoMail"
	"coderhub/shared/SQL"
	"coderhub/shared/Validator"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	Validator            *Validator.Validator
	UserRepository       repository.UserRepository
	ImageRelationService imagerelationservice.ImageRelationService
	ImageService         imageservice.ImageService
	RedisDB              CacheDB.RedisDB
	GoMail               GoMail.GoMailImpl
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := CacheDB.NewRedisDB(CacheDB.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:               c,
		Validator:            Validator.New(),
		UserRepository:       repository.NewUserRepositoryImpl(SQL.NewGorm(), redisDB),
		ImageRelationService: imagerelationservice.NewImageRelationService(zrpc.MustNewClient(c.ImageRelationService)),
		ImageService:         imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
		RedisDB:              redisDB,
		GoMail:               GoMail.NewGoMail(),
	}
}
