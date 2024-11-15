package logic

import (
	"coderhub/conf"
	"context"

	"coderhub/api/TechSphere/Articles/internal/svc"
	"coderhub/api/TechSphere/Articles/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthLogic) Health() (resp *types.HealthResp, err error) {
	return &types.HealthResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
	}, nil
}
