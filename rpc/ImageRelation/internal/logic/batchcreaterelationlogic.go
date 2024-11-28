package logic

import (
	"coderhub/model"
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchCreateRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchCreateRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchCreateRelationLogic {
	return &BatchCreateRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BatchCreateRelation 批量创建图片关系
func (l *BatchCreateRelationLogic) BatchCreateRelation(in *imageRelation.BatchCreateRelationRequest) (*imageRelation.BatchCreateRelationResponse, error) {
	imageRelations := make([]*model.ImageRelation, len(in.Relations))
	for i, relation := range in.Relations {
		imageRelations[i] = &model.ImageRelation{
			ImageID:    relation.ImageId,
			EntityID:   relation.EntityId,
			EntityType: relation.EntityType,
		}
	}
	err := l.svcCtx.ImageRelationRepository.BatchCreate(l.ctx, imageRelations)
	if err != nil {
		return nil, err
	}
	relations := make([]*imageRelation.ImageRelation, len(imageRelations))
	for i, relation := range imageRelations {
		relations[i] = &imageRelation.ImageRelation{
			ImageId:    relation.ImageID,
			EntityId:   relation.EntityID,
			EntityType: relation.EntityType,
		}
	}
	return &imageRelation.BatchCreateRelationResponse{
		Relations: relations,
	}, nil
}