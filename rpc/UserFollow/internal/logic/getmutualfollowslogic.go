package logic

import (
	"context"

	"coderhub/rpc/User/user"
	"coderhub/rpc/UserFollow/internal/svc"
	"coderhub/rpc/UserFollow/user_follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMutualFollowsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMutualFollowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMutualFollowsLogic {
	return &GetMutualFollowsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetMutualFollows 获取互相关注列表
func (l *GetMutualFollowsLogic) GetMutualFollows(in *user_follow.GetMutualFollowsReq) (*user_follow.GetMutualFollowsResp, error) {
	// 获取互相关注列表(ID)
	mutualFollows, err := l.svcCtx.UserFollowRepository.GetMutualFollows(in.UserId, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 获取互相关注列表(User)
	userIDs := make([]int64, 0, len(mutualFollows))
	for _, follow := range mutualFollows {
		userIDs = append(userIDs, follow.FollowedID)
	}
	users, err := l.svcCtx.UserService.BatchGetUserByID(l.ctx, &user.BatchGetUserByIDRequest{
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}
	// 转换为user_follow.UserInfo
	userInfos := make([]*user_follow.UserInfo, 0, len(users.UserInfos))
	for _, userInfo := range users.UserInfos {
		userInfos = append(userInfos, &user_follow.UserInfo{
			Id:       userInfo.UserId,
			Username: userInfo.UserName,
			Avatar:   userInfo.Avatar,
		})
	}

	return &user_follow.GetMutualFollowsResp{
		MutualFollows: userInfos,
		Total:         int64(len(mutualFollows)),
	}, nil
}
