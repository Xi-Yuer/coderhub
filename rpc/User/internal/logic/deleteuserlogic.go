package logic

import (
	"context"
	"errors"

	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	UserInfo, err := NewGetUserInfoLogic(l.ctx, l.svcCtx).GetUserInfo(&user.GetUserInfoRequest{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, err
	}
	if UserInfo.UserId == 0 {
		return nil, errors.New("用户不存在")
	}
	if err := l.svcCtx.UserRepository.DeleteUser(in.UserId); err != nil {
		return nil, err
	}

	return &user.DeleteUserResponse{
		Success: true,
	}, nil
}
