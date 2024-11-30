package logic

import (
	"context"
	"strconv"

	"coderhub/model"
	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetImagesByEntityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetImagesByEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetImagesByEntityLogic {
	return &BatchGetImagesByEntityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BatchGetImagesByEntity 批量获取图片，根据实体ID列表、实体类型列表获取
func (l *BatchGetImagesByEntityLogic) BatchGetImagesByEntity(in *imageRelation.BatchGetImagesByEntityRequest) (*imageRelation.BatchGetImagesByEntityResponse, error) {
	l.Logger.Infof("获取实体图片关系，实体IDs: %v, 实体类型: %s", in.EntityIds, in.EntityType)
	
	imageRelations, err := l.svcCtx.ImageRelationRepository.BatchGetImagesByEntity(l.ctx, in.EntityIds, in.EntityType)
	if err != nil {
		return nil, err
	}
	
	l.Logger.Infof("查询到的图片关系数量: %d", len(imageRelations))
	
	// 收集所有图片ID
	imageIds := make([]int64, len(imageRelations))
	for i, rel := range imageRelations {
			imageIds[i] = rel.ImageID
	}

	// 批量获取图片详细信息
	images, err := l.svcCtx.ImageRepository.BatchGetImagesByID(l.ctx, imageIds)
	if err != nil {
		return nil, err
	}

	// 创建图片ID到图片信息的映射
	imageMap := make(map[int64]*model.Image)
	for _, img := range images {
		imageMap[img.ID] = &img
	}

	// 转换类型并填充图片详细信息
	relations := make([]*imageRelation.ImageRelation, len(imageRelations))
	for i, rel := range imageRelations {
		img := imageMap[rel.ImageID]
		relations[i] = &imageRelation.ImageRelation{
			Id:           rel.ID,
			ImageId:      rel.ImageID,
			EntityId:     rel.EntityID,
			EntityType:   rel.EntityType,
			Url:          img.URL,
			ThumbnailUrl: img.ThumbnailURL,
			Sort:         rel.Sort,
			CreatedAt:    strconv.FormatInt(rel.CreatedAt.Unix(), 10),
		}
	}

	return &imageRelation.BatchGetImagesByEntityResponse{Relations: relations}, nil
}
