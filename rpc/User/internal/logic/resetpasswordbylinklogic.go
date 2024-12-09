package logic

import (
	"context"

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

// 通过链接重置密码
func (l *ResetPasswordByLinkLogic) ResetPasswordByLink(in *user.ResetPasswordByLinkRequest) (*user.ResetPasswordByLinkResponse, error) {
	// todo: add your logic here and delete this line

	return &user.ResetPasswordByLinkResponse{}, nil
}
