package logic

import (
	"context"

	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetLogic {
	return &BatchGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量获取图片信息
func (l *BatchGetLogic) BatchGet(in *image.BatchGetRequest) (*image.BatchGetResponse, error) {
	// todo: add your logic here and delete this line

	return &image.BatchGetResponse{}, nil
}
