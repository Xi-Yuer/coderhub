package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"
	"coderhub/conf"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 健康检查
func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLogic) Health() (resp *types.HealthResponse, err error) {
	return &types.HealthResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
	}, nil
}
