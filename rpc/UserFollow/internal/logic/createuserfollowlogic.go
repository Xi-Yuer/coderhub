package logic

import (
	"context"

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

// 创建用户关注关系
func (l *CreateUserFollowLogic) CreateUserFollow(in *user_follow.CreateUserFollowReq) (*user_follow.CreateUserFollowResp, error) {
	// todo: add your logic here and delete this line

	return &user_follow.CreateUserFollowResp{}, nil
}
