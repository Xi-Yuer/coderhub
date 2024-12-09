package logic

import (
	"context"

	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendResetPasswordLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送重置密码链接
func NewSendResetPasswordLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendResetPasswordLinkLogic {
	return &SendResetPasswordLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendResetPasswordLinkLogic) SendResetPasswordLink(req *types.SendResetPasswordLinkRequest) (resp *types.SendResetPasswordLinkResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
