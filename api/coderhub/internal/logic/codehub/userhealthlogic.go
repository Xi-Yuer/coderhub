package codehub

import (
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 健康检查
func NewUserHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHealthLogic {
	return &UserHealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHealthLogic) UserHealth() (resp *types.HealthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
