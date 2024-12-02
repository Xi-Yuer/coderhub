package svc

import (
	repository "coderhub/repository/Comment"
	"coderhub/rpc/Image/imageservice"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/TechSphere/Comment/internal/config"
	"coderhub/rpc/User/userservice"
	"coderhub/shared/CacheDB"
	"coderhub/shared/MQ"
	"coderhub/shared/SQL"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config               config.Config
	ImageRelationService imagerelationservice.ImageRelationService
	ImageService         imageservice.ImageService
	UserService          userservice.UserService
	CommentRepository    repository.CommentRepository
	MQ                   *rabbitmq.RabbitMQ
	Consumer             *rabbitmq.Consumer
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := CacheDB.NewRedisDB(CacheDB.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:               c,
		ImageRelationService: imagerelationservice.NewImageRelationService(zrpc.MustNewClient(c.ImageRelationService)),
		ImageService:         imageservice.NewImageService(zrpc.MustNewClient(c.ImageService)),
		UserService:          userservice.NewUserService(zrpc.MustNewClient(c.UserService)),
		CommentRepository:    repository.NewCommentRepository(SQL.NewGorm(), redisDB),
		MQ:                   rabbitmq.NewServiceContext(c.RabbitMQ).MQ,
		Consumer:             rabbitmq.NewServiceContext(c.RabbitMQ).Consumer,
	}
}
