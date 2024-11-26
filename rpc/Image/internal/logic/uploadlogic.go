package logic

import (
	"context"

	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传图片
func (l *UploadLogic) Upload(in *image.UploadRequest) (*image.ImageInfo, error) {
	// todo: add your logic here and delete this line

	return &image.ImageInfo{}, nil
}
