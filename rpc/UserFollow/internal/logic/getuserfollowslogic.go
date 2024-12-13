package logic

import (
	"context"

	"coderhub/rpc/User/user"
	"coderhub/rpc/UserFollow/internal/svc"
	"coderhub/rpc/UserFollow/user_follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowsLogic {
	return &GetUserFollowsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserFollows 获取用户关注列表
func (l *GetUserFollowsLogic) GetUserFollows(in *user_follow.GetUserFollowsReq) (*user_follow.GetUserFollowsResp, error) {
	// 获取关注列表
	userFollows, err := l.svcCtx.UserFollowRepository.GetUserFollows(in.FollowerId, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 构建用户IDs
	userIDs := make([]int64, 0, len(userFollows))
	for _, userFollow := range userFollows {
		userIDs = append(userIDs, userFollow.FollowedID)
	}
	// 批量获取用户信息
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

	return &user_follow.GetUserFollowsResp{
		UserFollows: userInfos,
		Total:       int64(len(userFollows)),
	}, nil
}
