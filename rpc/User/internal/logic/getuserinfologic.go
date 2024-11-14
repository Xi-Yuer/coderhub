package logic

import (
	"coderhub/model"
	"context"

	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	var User *model.User
	User, err := l.svcCtx.UserRepository.GetUserByID(in.UserId)
	if err != nil {
		return nil, err
	}

	return &user.GetUserInfoResponse{
		UserId:    User.ID,
		UserName:  User.UserName,
		Avatar:    User.Avatar.String,
		Email:     User.Email.String,
		Password:  User.Password,
		NickName:  User.NickName.String,
		IsAdmin:   User.IsAdmin,
		Status:    User.Status,
		CreatedAt: User.CreatedAt.Unix(),
		UpdatedAt: User.UpdatedAt.Unix(),
	}, nil
}
