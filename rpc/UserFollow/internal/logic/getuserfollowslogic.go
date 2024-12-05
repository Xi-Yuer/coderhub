package logic

import (
	"context"

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

// 获取用户关注列表
func (l *GetUserFollowsLogic) GetUserFollows(in *user_follow.GetUserFollowsReq) (*user_follow.GetUserFollowsResp, error) {
	// todo: add your logic here and delete this line

	return &user_follow.GetUserFollowsResp{}, nil
}
