package logic

import (
	"context"

	"coderhub/rpc/TechSphere/AcademicNavigator/academic_navigator"
	"coderhub/rpc/TechSphere/AcademicNavigator/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeAcademicNavigatorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeAcademicNavigatorLogic {
	return &LikeAcademicNavigatorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞学术导航
func (l *LikeAcademicNavigatorLogic) LikeAcademicNavigator(in *academic_navigator.LikeAcademicNavigatorRequest) (*academic_navigator.Response, error) {
	// todo: add your logic here and delete this line

	return &academic_navigator.Response{}, nil
}
