package logic

import (
	"context"

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

// 获取互相关注列表
func (l *GetMutualFollowsLogic) GetMutualFollows(in *user_follow.GetMutualFollowsReq) (*user_follow.GetMutualFollowsResp, error) {
	// todo: add your logic here and delete this line

	return &user_follow.GetMutualFollowsResp{}, nil
}
