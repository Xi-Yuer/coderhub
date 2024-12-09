package logic

import (
	"context"

	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"
	"coderhub/rpc/User/user"

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
	err = l.svcCtx.Validator.Email(req.Email).Check()
	if err != nil {
		return l.errorResp(err), nil
	}
	_, err = l.svcCtx.UserService.ResetPassword(l.ctx, &user.ResetPasswordRequest{
		Email: req.Email,
	})
	if err != nil {
		return l.errorResp(err), nil
	}

	return l.successResp(), nil
}

func (l *SendResetPasswordLinkLogic) successResp() *types.SendResetPasswordLinkResponse {
	return &types.SendResetPasswordLinkResponse{
		Response: types.Response{
			Code:    0,
			Message: "发送成功",
		},
		Data: true,
	}
}

func (l *SendResetPasswordLinkLogic) errorResp(err error) *types.SendResetPasswordLinkResponse {
	return &types.SendResetPasswordLinkResponse{
		Response: types.Response{
			Code:    1,
			Message: err.Error(),
		},
		Data: false,
	}
}
