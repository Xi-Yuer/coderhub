package logic

import (
	"context"

	"coderhub/api/TechSphere/AcademicNavigator/internal/svc"
	"coderhub/api/TechSphere/AcademicNavigator/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostAcademicNavigatorLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 点赞学术导航
func NewPostAcademicNavigatorLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostAcademicNavigatorLikeLogic {
	return &PostAcademicNavigatorLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostAcademicNavigatorLikeLogic) PostAcademicNavigatorLike(req *types.PostAcademicNavigatorLikeReq) (resp *types.PostAcademicNavigatorLikeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
