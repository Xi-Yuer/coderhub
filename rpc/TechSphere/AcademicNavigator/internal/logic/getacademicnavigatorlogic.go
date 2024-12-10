package logic

import (
	"context"

	"coderhub/model"
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
func (l *GetAcademicNavigatorLogic) GetAcademicNavigator(in *academic_navigator.GetAcademicNavigatorRequest) (*academic_navigator.GetAcademicNavigatorResponse, error) {
	academicNavigators, total, err := l.svcCtx.AcademicNavigatorRepository.GetAcademicNavigator(&model.AcademicNavigator{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, err
	}

	academicNavigatorList := make([]*academic_navigator.AcademicNavigator, 0, len(academicNavigators))
	academicNavigatorIDs := make([]int64, 0, len(academicNavigators))
	for _, academicNavigator := range academicNavigators {
		academicNavigatorIDs = append(academicNavigatorIDs, int64(academicNavigator.ID))
	}
	academicRelationLikeCountMap, err := l.svcCtx.AcademicRelationLikeRepository.BatchGetAcademicRelationLikeCount(l.ctx, academicNavigatorIDs)
	if err != nil {
		return nil, err
	}
	for _, academicNavigator := range academicNavigators {
		academicNavigatorList = append(academicNavigatorList, &academic_navigator.AcademicNavigator{
			Id:        int64(academicNavigator.ID),
			UserId:    academicNavigator.UserId,
			Content:   academicNavigator.Content,
			Education: academicNavigator.Education,
			Major:     academicNavigator.Major,
			School:    academicNavigator.School,
			WorkExp:   academicNavigator.WorkExp,
			LikeCount: academicRelationLikeCountMap[int64(academicNavigator.ID)],
			CreatedAt: academicNavigator.CreatedAt.Unix(),
			UpdatedAt: academicNavigator.UpdatedAt.Unix(),
		})
	}

	return &academic_navigator.GetAcademicNavigatorResponse{
		AcademicNavigator: academicNavigatorList,
		Total:             total,
	}, nil
}
