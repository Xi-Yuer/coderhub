package logic

import (
	"coderhub/api/user/internal/svc"
	"coderhub/api/user/internal/types"
	"coderhub/conf"
	"coderhub/rpc/user/user"
	"coderhub/shared/metaData"
	"coderhub/shared/validator"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordRequest) (resp *types.ChangePasswordResponse, err error) {
	if err := validator.New().Password(req.OldPassword).Password(req.NewPassword).Check(); err != nil {
		return &types.ChangePasswordResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
			Data: false,
		}, nil
	}

	ctx := metaData.SetUserMetaData(l.ctx) // 设置元数据
	response, err := l.svcCtx.UserService.ChangePassword(ctx, &user.ChangePasswordRequest{
		UserId:      req.UserId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})

	if err != nil {
		return &types.ChangePasswordResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
			Data: false,
		}, nil
	}

	// 返回成功响应
	return &types.ChangePasswordResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: response.Success,
	}, nil
}
