package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消关注
func NewDeleteUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserFollowLogic {
	return &DeleteUserFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserFollowLogic) DeleteUserFollow(req *types.DeleteUserFollowReq) (resp *types.DeleteUserFollowResp, err error) {
	// todo: add your logic here and delete this line

	return
}
