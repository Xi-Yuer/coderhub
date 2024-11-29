package logic

import (
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteRelationLogic {
	return &BatchDeleteRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除图片关系
func (l *BatchDeleteRelationLogic) BatchDeleteRelation(in *imageRelation.BatchDeleteRelationRequest) (*imageRelation.BatchDeleteRelationResponse, error) {
	// todo: add your logic here and delete this line

	return &imageRelation.BatchDeleteRelationResponse{}, nil
}
