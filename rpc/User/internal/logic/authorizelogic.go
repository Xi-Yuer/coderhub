package logic

import (
	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"
	"coderhub/shared/BcryptUtil"
	"coderhub/shared/JWT"
	"coderhub/shared/Validator"
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
	if err := Validator.New().Username(in.Username).Password(in.Password).Check(); err != nil {
		return nil, err
	}

	UserInfo, err := NewGetUserInfoByUsernameLogic(l.ctx, l.svcCtx).GetUserInfoByUsername(&user.GetUserInfoByUsernameRequest{Username: in.Username})
	if err != nil {
		return nil, err
	}

	if !BcryptUtil.CompareHashAndPassword(UserInfo.Password, in.Password) {
		return nil, errors.New("密码错误")
	}

	if authorization, err := JWT.GenerateAuthorization(UserInfo.UserId); err != nil {
		return nil, err
	} else {
		return &user.AuthorizeResponse{
			Token: authorization,
		}, nil
	}
}
