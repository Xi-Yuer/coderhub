package logic

import (
	"context"

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

// 获取用户粉丝列表
func (l *GetUserFansLogic) GetUserFans(in *user_follow.GetUserFansReq) (*user_follow.GetUserFansResp, error) {
	// todo: add your logic here and delete this line

	return &user_follow.GetUserFansResp{}, nil
}
