package logic

import (
	"context"

	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"
	"coderhub/rpc/User/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordByLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过链接重置密码
func NewResetPasswordByLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordByLinkLogic {
	return &ResetPasswordByLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordByLinkLogic) ResetPasswordByLink(req *types.ResetPasswordByLinkRequest) (resp *types.ResetPasswordByLinkResponse, err error) {
	err = l.svcCtx.Validator.Email(req.Email).Password(req.Password).ConfirmPassword(req.Password, req.ConfirmPassword).Token(req.Token).Check()
	if err != nil {
		return l.errorResp(err), nil
	}
	_, err = l.svcCtx.UserService.ResetPasswordByLink(l.ctx, &user.ResetPasswordByLinkRequest{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		Token:           req.Token,
	})
	if err != nil {
		return l.errorResp(err), nil
	}

	return l.successResp(), nil
}

func (l *ResetPasswordByLinkLogic) successResp() *types.ResetPasswordByLinkResponse {
	return &types.ResetPasswordByLinkResponse{
		Response: types.Response{
			Code:    0,
			Message: "重置密码成功",
		},
		Data: true,
	}
}

func (l *ResetPasswordByLinkLogic) errorResp(err error) *types.ResetPasswordByLinkResponse {
	return &types.ResetPasswordByLinkResponse{
		Response: types.Response{
			Code:    1,
			Message: err.Error(),
		},
		Data: false,
	}
}
