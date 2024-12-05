package logic

import (
	"context"

	"coderhub/rpc/UserFollow/internal/svc"
	"coderhub/rpc/UserFollow/user_follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserFollowLogic {
	return &DeleteUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户关注关系
func (l *DeleteUserFollowLogic) DeleteUserFollow(in *user_follow.DeleteUserFollowReq) (*user_follow.DeleteUserFollowResp, error) {
	// todo: add your logic here and delete this line

	return &user_follow.DeleteUserFollowResp{}, nil
}
