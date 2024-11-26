package logic

import (
	"context"

	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByUserLogic {
	return &ListByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户图片列表
func (l *ListByUserLogic) ListByUser(in *image.ListByUserRequest) (*image.ListByUserResponse, error) {
	// todo: add your logic here and delete this line

	return &image.ListByUserResponse{}, nil
}
