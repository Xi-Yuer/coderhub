package logic

import (
	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"
	"coderhub/conf"
	"coderhub/rpc/User/user"
	"coderhub/shared/Validator"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserRequest) (resp *types.CreateUserResponse, err error) {
	if err := Validator.New().Username(req.Username).Password(req.PasswordHash).Check(); err != nil {
		return &types.CreateUserResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
			Data: 0,
		}, nil
	}

	var createUserResponse *user.CreateUserResponse

	if createUserResponse, err = l.svcCtx.UserService.CreateUser(
		l.ctx,
		&user.CreateUserRequest{
			Username:     req.Username,
			PasswordHash: req.PasswordHash,
		},
	); err != nil {
		return &types.CreateUserResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
			Data: 0,
		}, nil
	}

	return &types.CreateUserResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: createUserResponse.UserId,
	}, nil
}
