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
	if err := l.svcCtx.ImageRelationRepository.BatchDelete(l.ctx, in.Ids); err != nil {
		return nil, err
	}

	return &imageRelation.BatchDeleteRelationResponse{
		Success: true,
	}, nil
}
