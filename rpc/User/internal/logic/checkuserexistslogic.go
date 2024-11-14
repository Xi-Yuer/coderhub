package logic

import (
	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserExistsLogic {
	return &CheckUserExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserExistsLogic) CheckUserExists(in *user.CheckUserExistsRequest) (*user.CheckUserExistsResponse, error) {
	UserInfo, err := NewGetUserInfoByUsernameLogic(l.ctx, l.svcCtx).GetUserInfoByUsername(&user.GetUserInfoByUsernameRequest{Username: in.Username})

	if UserInfo == nil || err != nil {
		return &user.CheckUserExistsResponse{
			Exists: false,
		}, nil
	}
	return &user.CheckUserExistsResponse{
		Exists: true,
	}, nil
}
