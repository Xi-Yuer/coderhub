package logic

import (
	"context"

	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoByUsernameLogic {
	return &GetUserInfoByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoByUsernameLogic) GetUserInfoByUsername(in *user.GetUserInfoByUsernameRequest) (*user.GetUserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserInfoResponse{
		UserId:    0,
		Username:  "",
		Avatar:    "",
		Email:     "",
		Nickname:  "",
		IsAdmin:   "",
		Status:    false,
		CreatedAt: 0,
		UpdatedAt: 0,
	}, nil
}
