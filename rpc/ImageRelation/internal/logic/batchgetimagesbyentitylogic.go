package logic

import (
	"context"

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

// 批量获取图片关联，根据实体ID列表、实体类型列表获取
func (l *BatchGetImagesByEntityLogic) BatchGetImagesByEntity(in *imageRelation.BatchGetImagesByEntityRequest) (*imageRelation.BatchGetImagesByEntityResponse, error) {
	// todo: add your logic here and delete this line

	return &imageRelation.BatchGetImagesByEntityResponse{}, nil
}
