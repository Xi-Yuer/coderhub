package logic

import (
	"coderhub/api/user/internal/svc"
	"coderhub/api/user/internal/types"
	"coderhub/conf"
	"coderhub/rpc/user/user"
	"coderhub/shared/validator"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticateUserLogic {
	return &AuthenticateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticateUserLogic) AuthenticateUser(req *types.AuthenticateUserRequest) (resp *types.AuthenticateUserResponse, err error) {
	if err := validator.New().Username(req.Username).Password(req.Password).Check(); err != nil {
		return &types.AuthenticateUserResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
			Data: "",
		}, nil
	}

	var authorize *user.AuthorizeResponse

	if authorize, err = l.svcCtx.UserService.Authorize(l.ctx, &user.AuthorizeRequest{
		Username: req.Username,
		Password: req.Password,
	}); err != nil {
		return &types.AuthenticateUserResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
			Data: "",
		}, nil
	}

	return &types.AuthenticateUserResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: authorize.Token,
	}, nil
}
