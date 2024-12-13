package logic

import (
	"coderhub/conf"
	"coderhub/shared/utils"
	"context"

	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordRequest) (resp *types.ResetPasswordResponse, err error) {
	if err := utils.NewValidator().Password(req.Email).Password(req.NewPassword).Check(); err != nil {
		return &types.ResetPasswordResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
			Data: false,
		}, nil
	}

	_, err = utils.GetUserID(l.ctx)
	if err != nil {
		return &types.ResetPasswordResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
		}, nil
	}
	_ = utils.SetUserMetaData(l.ctx) // 设置元数据

	return
}
