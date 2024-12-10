package logic

import (
	"context"

	"coderhub/api/TechSphere/AcademicNavigator/internal/svc"
	"coderhub/api/TechSphere/AcademicNavigator/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAcademicNavigatorLikeCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新学术导航点赞数
func NewUpdateAcademicNavigatorLikeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAcademicNavigatorLikeCountLogic {
	return &UpdateAcademicNavigatorLikeCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAcademicNavigatorLikeCountLogic) UpdateAcademicNavigatorLikeCount(req *types.UpdateAcademicNavigatorLikeCountReq) (resp *types.UpdateAcademicNavigatorLikeCountResp, err error) {
	// todo: add your logic here and delete this line

	return
}
