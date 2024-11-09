package logic

import (
	"coderhub/model"
	"context"
	"errors"

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
	var User model.User

	if err := l.svcCtx.SqlDB.First(&User, "user_name = ?", in.Username).Error; err != nil {
		return nil, err
	}
	if User.ID == 0 {
		return nil, errors.New("用户不存在")
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
