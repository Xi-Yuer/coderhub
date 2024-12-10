package logic

import (
	"context"

	"coderhub/rpc/TechSphere/AcademicNavigator/academic_navigator"
	"coderhub/rpc/TechSphere/AcademicNavigator/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeAcademicNavigatorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLikeAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeAcademicNavigatorLogic {
	return &CancelLikeAcademicNavigatorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消点赞学术导航
func (l *CancelLikeAcademicNavigatorLogic) CancelLikeAcademicNavigator(in *academic_navigator.CancelLikeAcademicNavigatorRequest) (*academic_navigator.Response, error) {
	// todo: add your logic here and delete this line

	return &academic_navigator.Response{}, nil
}
