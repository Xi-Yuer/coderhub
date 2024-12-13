package logic

import (
	"coderhub/shared/security"
	"context"
	"errors"
	"fmt"

	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordByLinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordByLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordByLinkLogic {
	return &ResetPasswordByLinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ResetPasswordByLink 通过链接重置密码
func (l *ResetPasswordByLinkLogic) ResetPasswordByLink(in *user.ResetPasswordByLinkRequest) (*user.ResetPasswordByLinkResponse, error) {
	err := l.svcCtx.Validator.Email(in.Email).Password(in.Password).ConfirmPassword(in.Password, in.ConfirmPassword).Token(in.Token).Check()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// 1. 获取redis中的token
	token, err := l.svcCtx.RedisDB.Get(fmt.Sprintf("reset_password:%s", in.Email))
	if err != nil {
		return nil, err
	}
	// 2. 判断token是否正确
	if token != in.Token {
		return nil, fmt.Errorf("token不正确")
	}
	// 3. 删除redis中的token
	err = l.svcCtx.RedisDB.Del(fmt.Sprintf("reset_password:%s", in.Email))
	if err != nil {
		return nil, err
	}
	// 4. 更新用户密码
	password, err := security.PasswordHash(in.Password)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.UserRepository.ResetPassword(in.Email, password)
	if err != nil {
		return nil, err
	}

	return &user.ResetPasswordByLinkResponse{
		Success: true,
	}, nil
}
