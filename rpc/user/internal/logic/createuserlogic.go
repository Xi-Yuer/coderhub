package logic

import (
	"coderhub/model"
	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"
	"coderhub/shared/bcryptUtil"
	"coderhub/shared/snowFlake"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	exists, _ := NewCheckUserExistsLogic(l.ctx, l.svcCtx).CheckUserExists(&user.CheckUserExistsRequest{Username: in.Username})
	if exists.Exists {
		return nil, errors.New("用户已存在")
	}

	ID := snowFlake.GenID()
	// 密码加密
	Password, _ := bcryptUtil.PasswordHash(in.PasswordHash)
	if tx := l.svcCtx.SqlDB.Create(&model.User{
		ID:       ID,
		UserName: in.Username,
		Password: Password,
	}); tx.Error != nil {
		return nil, tx.Error
	}

	return &user.CreateUserResponse{
		UserId: ID,
	}, nil
}
