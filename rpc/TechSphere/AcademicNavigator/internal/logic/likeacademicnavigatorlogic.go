package logic

import (
	"context"

	"coderhub/model"
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
	err := l.svcCtx.AcademicRelationLikeRepository.AddAcademicRelationLike(l.ctx, &model.AcademicRelationLike{
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
