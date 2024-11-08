package logic

import (
	"coderhub/model"
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
	var User model.User

	if tx := l.svcCtx.SqlDB.First(&User, "user_name = ?", in.Username); tx.Error != nil {
		return nil, tx.Error
	}
	if User.UserName == "" {
		return nil, errors.New("用户不存在")
	}

	if !bcryptUtil.CompareHashAndPassword(User.Password, in.Password) {
		return nil, errors.New("密码错误")
	}

	if authorization, err := token.GenerateAuthorization(User.ID); err != nil {
		return nil, err
	} else {
		return &user.AuthorizeResponse{
			Token: authorization,
		}, nil
	}
}
