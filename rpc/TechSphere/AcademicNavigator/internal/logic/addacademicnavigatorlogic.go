package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &academic_navigator.Response{}, nil
}
