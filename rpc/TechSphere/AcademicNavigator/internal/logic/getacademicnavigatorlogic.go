package logic

import (
	"context"

	"coderhub/rpc/TechSphere/AcademicNavigator/academic_navigator"
	"coderhub/rpc/TechSphere/AcademicNavigator/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAcademicNavigatorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAcademicNavigatorLogic {
	return &GetAcademicNavigatorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取学术导航
func (l *GetAcademicNavigatorLogic) GetAcademicNavigator(in *academic_navigator.GetAcademicNavigatorRequest) (*academic_navigator.Response, error) {
	// todo: add your logic here and delete this line

	return &academic_navigator.Response{}, nil
}
