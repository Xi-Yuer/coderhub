package logic

import (
	"context"

	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除图片
func (l *DeleteLogic) Delete(in *image.DeleteRequest) (*image.DeleteResponse, error) {
	err := l.svcCtx.ImageRepository.Delete(l.ctx, in.ImageId)
	if err != nil {
		return nil, err
	}
	return &image.DeleteResponse{
		Success: true,
	}, nil
}
