package svc

import (
	repository "coderhub/repository/Comment"
	"coderhub/rpc/TechSphere/Comment/internal/config"
	"coderhub/shared/CacheDB"
	"coderhub/shared/SQL"
)

type ServiceContext struct {
	Config            config.Config
	CommentRepository repository.CommentRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := CacheDB.NewRedisDB(CacheDB.DefaultConfig())
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:            c,
		CommentRepository: repository.NewCommentRepository(SQL.NewGorm(), redisDB),
	}
}
