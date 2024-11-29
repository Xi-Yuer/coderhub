package logic

import (
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteByEntityIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteByEntityIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteByEntityIDLogic {
	return &DeleteByEntityIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据实体ID、实体类型删除图片关系
func (l *DeleteByEntityIDLogic) DeleteByEntityID(in *imageRelation.DeleteByEntityIDRequest) (*imageRelation.DeleteByEntityIDResponse, error) {
	// todo: add your logic here and delete this line

	return &imageRelation.DeleteByEntityIDResponse{}, nil
}
