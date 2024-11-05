package logic

import (
	"context"

	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthenticateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateUserLogic {
	return &AuthenticateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AuthenticateUser 验证用户登录
func (l *AuthenticateUserLogic) AuthenticateUser(in *user.AuthenticateUserRequest) (*user.AuthenticateUserResponse, error) {
	return &user.AuthenticateUserResponse{
		Authenticated: true,
	}, nil
}
