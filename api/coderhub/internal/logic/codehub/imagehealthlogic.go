package codehub

import (
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageHealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 健康检查
func NewImageHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageHealthLogic {
	return &ImageHealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageHealthLogic) ImageHealth() (resp *types.HealthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
