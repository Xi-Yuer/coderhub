package logic

import (
	"context"

	"coderhub/model"
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

// DeleteUserFollow 删除用户关注关系
func (l *DeleteUserFollowLogic) DeleteUserFollow(in *user_follow.DeleteUserFollowReq) (*user_follow.DeleteUserFollowResp, error) {
	err := l.svcCtx.UserFollowRepository.DeleteUserFollow(&model.UserFollow{
		FollowerID: in.FollowerId,
		FollowedID: in.FollowedId,
	})
	if err != nil {
		return nil, err
	}

	return &user_follow.DeleteUserFollowResp{
		Success: true,
	}, nil
}
