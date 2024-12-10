package logic

import (
	"context"

	"coderhub/rpc/TechSphere/AcademicNavigator/academic_navigator"
	"coderhub/rpc/TechSphere/AcademicNavigator/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAcademicNavigatorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAcademicNavigatorLogic {
	return &DeleteAcademicNavigatorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除学术导航
func (l *DeleteAcademicNavigatorLogic) DeleteAcademicNavigator(in *academic_navigator.DeleteAcademicNavigatorRequest) (*academic_navigator.Response, error) {
	// todo: add your logic here and delete this line

	return &academic_navigator.Response{}, nil
}
