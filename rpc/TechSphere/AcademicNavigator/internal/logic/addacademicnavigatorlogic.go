package logic

import (
	"context"

	"coderhub/model"
	"coderhub/rpc/TechSphere/AcademicNavigator/academic_navigator"
	"coderhub/rpc/TechSphere/AcademicNavigator/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAcademicNavigatorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAcademicNavigatorLogic {
	return &AddAcademicNavigatorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 新增学术导航
func (l *AddAcademicNavigatorLogic) AddAcademicNavigator(in *academic_navigator.AddAcademicNavigatorRequest) (*academic_navigator.Response, error) {
	err := l.svcCtx.AcademicNavigatorRepository.AddAcademicNavigator(&model.AcademicNavigator{
		UserId:    in.UserId,
		Content:   in.Content,
		Education: in.Education,
		Major:     in.Major,
		School:    in.School,
		WorkExp:   in.WorkExp,
	})
	if err != nil {
		return nil, err
	}

	return &academic_navigator.Response{
		Success: true,
	}, nil
}
