package logic

import (
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

// 批量创建图片关系
func (l *BatchCreateRelationLogic) BatchCreateRelation(in *imageRelation.BatchCreateRelationRequest) (*imageRelation.BatchCreateRelationResponse, error) {
	// todo: add your logic here and delete this line

	return &imageRelation.BatchCreateRelationResponse{}, nil
}
