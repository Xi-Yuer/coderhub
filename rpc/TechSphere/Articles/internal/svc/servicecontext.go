package svc

import (
	repository "coderhub/repository/Article"
	"coderhub/rpc/TechSphere/Articles/internal/config"
	"coderhub/shared/CacheDB"
	"coderhub/shared/SQL"
)

type ServiceContext struct {
	Config            config.Config
	ArticleRepository repository.ArticleRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisDB, err := CacheDB.NewRedisDB(CacheDB.DefaultConfig())
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:            c,
		ArticleRepository: repository.NewArticleRepositoryImpl(SQL.NewGorm(), redisDB),
	}
}
