package logic

import (
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEntitiesByImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEntitiesByImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEntitiesByImageLogic {
	return &GetEntitiesByImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetEntitiesByImage 获取图片关联的实体列表
func (l *GetEntitiesByImageLogic) GetEntitiesByImage(in *imageRelation.GetEntitiesByImageRequest) (*imageRelation.GetEntitiesByImageResponse, error) {
	imageRelations, err := l.svcCtx.ImageRelationRepository.ListByImageID(l.ctx, in.ImageId)
	if err != nil {
		return nil, err
	}

	entities := make([]*imageRelation.ImageRelation, len(imageRelations))
	for i, relation := range imageRelations {
		entities[i] = &imageRelation.ImageRelation{
			Id:         relation.ID,
			ImageId:    relation.ImageID,
			EntityId:   relation.EntityID,
			EntityType: relation.EntityType,
		}
	}
	entityInfos := make([]*imageRelation.EntityInfo, 0)
	for _, relation := range imageRelations {
		entityInfos = append(entityInfos, &imageRelation.EntityInfo{
			EntityType: relation.EntityType,
			EntityId:   relation.EntityID,
			CreatedAt:  relation.CreatedAt.Unix(),
		})
	}
	return &imageRelation.GetEntitiesByImageResponse{
		Entities: entityInfos,
	}, nil
}
