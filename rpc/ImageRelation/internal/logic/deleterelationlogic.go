package logic

import (
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRelationLogic {
	return &DeleteRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除图片关系
func (l *DeleteRelationLogic) DeleteRelation(in *imageRelation.DeleteRelationRequest) (*imageRelation.DeleteRelationResponse, error) {
	// todo: add your logic here and delete this line

	return &imageRelation.DeleteRelationResponse{}, nil
}
