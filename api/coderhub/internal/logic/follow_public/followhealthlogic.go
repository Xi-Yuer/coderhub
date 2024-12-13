package follow_public

import (
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowHealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 健康检查
func NewFollowHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowHealthLogic {
	return &FollowHealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowHealthLogic) FollowHealth() (resp *types.HealthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
