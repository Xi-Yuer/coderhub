package logic

import (
	"context"

	"coderhub/model"
	"coderhub/rpc/UserFollow/internal/svc"
	"coderhub/rpc/UserFollow/user_follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserFollowLogic {
	return &CreateUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateUserFollow 创建用户关注关系
func (l *CreateUserFollowLogic) CreateUserFollow(in *user_follow.CreateUserFollowReq) (*user_follow.CreateUserFollowResp, error) {
	// 是否已经关注
	exist, err := l.svcCtx.UserFollowRepository.IsUserFollowed(in.FollowerId, in.FollowedId)
	if err != nil {
		return nil, err
	}
	if exist {
		// 取消关注
		err = l.svcCtx.UserFollowRepository.DeleteUserFollow(&model.UserFollow{
			FollowerID: in.FollowerId,
			FollowedID: in.FollowedId,
		})
		if err != nil {
			return nil, err
		}
		return &user_follow.CreateUserFollowResp{
			Success: true,
		}, nil
	}

	if err := l.svcCtx.UserFollowRepository.CreateUserFollow(&model.UserFollow{
		FollowerID: in.FollowerId,
		FollowedID: in.FollowedId,
	}); err != nil {
		return nil, err
	}

	return &user_follow.CreateUserFollowResp{
		Success: true,
	}, nil
}
