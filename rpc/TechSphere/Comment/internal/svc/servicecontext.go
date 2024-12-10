package svc

import (
	"coderhub/repository"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/TechSphere/Comment/internal/config"
	"coderhub/rpc/User/userservice"
	"coderhub/shared/storage"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                        config.Config
	ImageRelationService          imagerelationservice.ImageRelationService
	ImageService                  imageservice.ImageService
	UserService                   userservice.UserService
	CommentRepository             repository.CommentRepository
	CommentRelationLikeRepository repository.CommentRelationLikeRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := storage.NewRedisDB(storage.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:                        c,
		ImageRelationService:          imagerelationservice.NewImageRelationService(zrpc.MustNewClient(c.ImageRelationService)),
		ImageService:                  imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
		UserService:                   userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
		CommentRepository:             repository.NewCommentRepository(storage.NewGorm(), redisDB),
		CommentRelationLikeRepository: repository.NewCommentRelationLikeRepository(storage.NewGorm(), redisDB),
	}
}
