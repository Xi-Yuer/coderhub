package logic

import (
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImagesByEntityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetImagesByEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImagesByEntityLogic {
	return &GetImagesByEntityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取实体关联的图片列表
func (l *GetImagesByEntityLogic) GetImagesByEntity(in *imageRelation.GetImagesByEntityRequest) (*imageRelation.GetImagesByEntityResponse, error) {
	// todo: add your logic here and delete this line

	return &imageRelation.GetImagesByEntityResponse{}, nil
}
