package logic

import (
	"coderhub/rpc/user/user"
	"context"

	"coderhub/api/user/internal/svc"
	"coderhub/api/user/internal/types"

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
				Code:    -1,
				Message: err.Error(),
			},
			Data: 0,
		}, nil
	}

	return &types.CreateUserResponse{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		Data: createUserResponse.UserId,
	}, nil
}
