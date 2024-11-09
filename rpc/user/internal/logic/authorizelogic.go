package logic

import (
	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"
	"coderhub/shared/bcryptUtil"
	"coderhub/shared/token"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthorizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthorizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthorizeLogic {
	return &AuthorizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthorizeLogic) Authorize(in *user.AuthorizeRequest) (*user.AuthorizeResponse, error) {
	UserInfo, err := NewGetUserInfoByUsernameLogic(l.ctx, l.svcCtx).GetUserInfoByUsername(&user.GetUserInfoByUsernameRequest{Username: in.Username})
	if err != nil {
		return nil, err
	}

	if !bcryptUtil.CompareHashAndPassword(UserInfo.Password, in.Password) {
		return nil, errors.New("密码错误")
	}

	if authorization, err := token.GenerateAuthorization(UserInfo.UserId); err != nil {
		return nil, err
	} else {
		return &user.AuthorizeResponse{
			Token: authorization,
		}, nil
	}
}
