package svc

import (
	"coderhub/api/TechSphere/AcademicNavigator/internal/config"
	"coderhub/rpc/TechSphere/AcademicNavigator/academicnavigatorservice"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                   config.Config
	AcademicNavigatorService academicnavigatorservice.AcademicNavigatorService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                   c,
		AcademicNavigatorService: academicnavigatorservice.NewAcademicNavigatorService(zrpc.MustNewClient(c.AcademicNavigatorService)),
	}
}
