package logic

import (
	"context"

	"coderhub/rpc/User/user"
	"coderhub/rpc/UserFollow/internal/svc"
	"coderhub/rpc/UserFollow/user_follow"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFansLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFansLogic {
	return &GetUserFansLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserFans 获取用户粉丝列表
func (l *GetUserFansLogic) GetUserFans(in *user_follow.GetUserFansReq) (*user_follow.GetUserFansResp, error) {
	// 获取粉丝列表
	userFollows, err := l.svcCtx.UserFollowRepository.GetUserFans(in.FollowedId, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 构建粉丝IDs
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
	for _, user := range users.UserInfos {
		userInfos = append(userInfos, &user_follow.UserInfo{
			Id:       user.UserId,
			Username: user.UserName,
			Avatar:   user.Avatar,
		})
	}

	return &user_follow.GetUserFansResp{
		UserFans: userInfos,
		Total:    int64(len(userFollows)),
	}, nil
}
