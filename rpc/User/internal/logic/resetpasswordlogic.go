package logic

import (
	"context"

	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"
	"coderhub/shared/Validator"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetPasswordLogic) ResetPassword(in *user.ResetPasswordRequest) (*user.ResetPasswordResponse, error) {
	if err := Validator.New().Password(in.NewPassword).Check(); err != nil {
		return nil, err
	}

	return &user.ResetPasswordResponse{}, nil
}
