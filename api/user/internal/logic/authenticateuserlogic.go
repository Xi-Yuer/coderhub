package logic

import (
	"coderhub/rpc/user/user"
	"context"

	"coderhub/api/user/internal/svc"
	"coderhub/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateUserLogic {
	return &AuthenticateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticateUserLogic) AuthenticateUser(req *types.AuthenticateUserRequest) (resp *types.AuthenticateUserResponse, err error) {
	exists, err := l.svcCtx.UserService.CheckUserExists(l.ctx, &user.CheckUserExistsRequest{
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	if !exists.Exists {
		return nil, err
	}

	return
}
