package logic

import (
	"context"

	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取图片信息
func (l *GetLogic) Get(in *image.GetRequest) (*image.ImageInfo, error) {
	// todo: add your logic here and delete this line

	return &image.ImageInfo{}, nil
}
