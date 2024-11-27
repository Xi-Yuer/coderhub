package logic

import (
	"context"

	"coderhub/model"
	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRelationLogic {
	return &CreateRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建图片关系
func (l *CreateRelationLogic) CreateRelation(in *imageRelation.CreateRelationRequest) (*imageRelation.CreateRelationResponse, error) {
	err := l.svcCtx.ImageRelationRepository.Create(l.ctx, &model.ImageRelation{
		ImageID:    in.ImageId,
		EntityID:   in.EntityId,
		EntityType: in.EntityType,
	})
	if err != nil {
		return nil, err
	}

	return &imageRelation.CreateRelationResponse{
		Relation: &imageRelation.ImageRelation{
			ImageId:    in.ImageId,
			EntityId:   in.EntityId,
			EntityType: in.EntityType,
			Sort:       0,
		},
	}, nil
}
