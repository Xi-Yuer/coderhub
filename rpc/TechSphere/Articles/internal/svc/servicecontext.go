package svc

import (
	repository "coderhub/repository/Article"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/TechSphere/Articles/internal/config"
	"coderhub/shared/CacheDB"
	"coderhub/shared/SQL"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	ArticleRepository    repository.ArticleRepository
	ImageRelationService imagerelationservice.ImageRelationService
	ImageService         imageservice.ImageService
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := CacheDB.NewRedisDB(CacheDB.DefaultConfig())
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:            c,
		ArticleRepository: repository.NewArticleRepositoryImpl(SQL.NewGorm(), redisDB),
		ImageRelationService: imagerelationservice.NewImageRelationService(zrpc.MustNewClient(c.ImageRelationService)),
		ImageService:         imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
	}
}
