package svc

import (
	"coderhub/repository"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/TechSphere/Articles/internal/config"
	"coderhub/shared/storage"
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
	redisDB, err := storage.NewRedisDB(storage.DefaultConfig())
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:                         c,
		ImageRelationService:           imagerelationservice.NewImageRelationService(zrpc.MustNewClient(c.ImageRelationService)),
		ImageService:                   imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
		ArticleRepository:              repository.NewArticleRepositoryImpl(storage.NewGorm(), redisDB),
		ArticlesRelationLikeRepository: repository.NewArticlesRelationLikeRepository(storage.NewGorm(), redisDB),
		ArticlePVRepository:            repository.NewArticlePVRepositoryImpl(storage.NewGorm(), redisDB),
		CommentRepository:              repository.NewCommentRepository(storage.NewGorm(), redisDB),
	}
}
