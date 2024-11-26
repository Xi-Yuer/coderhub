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

// 获取图片关联的实体列表
func (l *GetEntitiesByImageLogic) GetEntitiesByImage(in *imageRelation.GetEntitiesByImageRequest) (*imageRelation.GetEntitiesByImageResponse, error) {
	// todo: add your logic here and delete this line

	return &imageRelation.GetEntitiesByImageResponse{}, nil
}
