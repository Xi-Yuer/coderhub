package svc

import (
	"coderhub/repository"
	"coderhub/rpc/TechSphere/AcademicNavigator/internal/config"
	"coderhub/shared/storage"

	"github.com/elastic/go-elasticsearch/v8"
)

type ServiceContext struct {
	Config                         config.Config
	AcademicNavigatorRepository    repository.AcademicNavigatorRepository
	AcademicRelationLikeRepository repository.AcademicRelationLikeRepository
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AcademicNavigatorRepository: repository.NewAcademicNavigatorRepositoryImpl(storage.NewGorm(), &elasticsearch.Config{
			Addresses: []string{"http://elasticsearch:9200"},
			Username:  "elastic",
			Password:  "2214380963Wx!!",
		}),
		AcademicRelationLikeRepository: repository.NewAcademicRelationLikeRepositoryImpl(storage.NewGorm()),
	}
}
