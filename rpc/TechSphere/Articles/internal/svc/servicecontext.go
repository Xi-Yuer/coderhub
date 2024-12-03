package svc

import (
	"coderhub/repository"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/TechSphere/Articles/internal/config"
	"coderhub/shared/CacheDB"
	"coderhub/shared/SQL"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                         config.Config
	ImageRelationService           imagerelationservice.ImageRelationService
	ImageService                   imageservice.ImageService
	ArticleRepository              repository.ArticleRepository
	ArticlesRelationLikeRepository repository.ArticlesRelationLikeRepository
	ArticlePVRepository            repository.ArticlePVRepository
	CommentRepository              repository.CommentRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := CacheDB.NewRedisDB(CacheDB.DefaultConfig())
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:                         c,
		ImageRelationService:           imagerelationservice.NewImageRelationService(zrpc.MustNewClient(c.ImageRelationService)),
		ImageService:                   imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
		ArticleRepository:              repository.NewArticleRepositoryImpl(SQL.NewGorm(), redisDB),
		ArticlesRelationLikeRepository: repository.NewArticlesRelationLikeRepository(SQL.NewGorm(), redisDB),
		ArticlePVRepository:            repository.NewArticlePVRepositoryImpl(SQL.NewGorm(), redisDB),
		CommentRepository:              repository.NewCommentRepository(SQL.NewGorm(), redisDB),
	}
}
