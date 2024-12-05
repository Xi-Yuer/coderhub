package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsUserFollowedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 检查是否关注
func NewIsUserFollowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserFollowedLogic {
	return &IsUserFollowedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsUserFollowedLogic) IsUserFollowed(req *types.IsUserFollowedReq) (resp *types.IsUserFollowedResp, err error) {
	// todo: add your logic here and delete this line

	return
}
