package logic

import (
	"coderhub/rpc/user/user"
	"coderhub/shared/token"
	"context"

	"coderhub/api/user/internal/svc"
	"coderhub/api/user/internal/types"

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
	exists, err := l.svcCtx.UserService.CheckUserExists(l.ctx, &user.CheckUserExistsRequest{
		Username: req.Username,
	})
	if err != nil {
		return &types.AuthenticateUserResponse{
			Response: types.Response{
				Code:    0,
				Message: "fail",
			},
			Data: err.Error()}, nil
	}
	if !exists.Exists {
		return &types.AuthenticateUserResponse{
			Response: types.Response{
				Code:    0,
				Message: "fail",
			},
			Data: "用户不存在",
		}, nil
	}

	UserInfo, err := l.svcCtx.UserService.GetUserInfoByUsername(l.ctx, &user.GetUserInfoByUsernameRequest{Username: req.Username})
	if err != nil {
		return &types.AuthenticateUserResponse{
			Response: types.Response{
				Code:    0,
				Message: "fail",
			},
			Data: err.Error()}, nil
	}

	authorization, err := token.GenerateAuthorization(UserInfo.UserId)
	if err != nil {
		return &types.AuthenticateUserResponse{
			Response: types.Response{
				Code:    0,
				Message: "fail",
			},
			Data: err.Error()}, nil
	}
	return &types.AuthenticateUserResponse{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		Data: authorization,
	}, nil
}
