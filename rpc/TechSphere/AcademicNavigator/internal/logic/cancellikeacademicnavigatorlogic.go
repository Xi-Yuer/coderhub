package logic

import (
	"context"

	"coderhub/model"
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
	err := l.svcCtx.AcademicRelationLikeRepository.DeleteAcademicRelationLike(l.ctx, &model.AcademicRelationLike{
		AcademicNavigatorID: in.Id,
		UserID:              in.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &academic_navigator.Response{
		Success: true,
	}, nil
}
