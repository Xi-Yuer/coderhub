package logic

import (
	"coderhub/model"
	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"
	"coderhub/shared/bcryptUtil"
	"coderhub/shared/snowFlake"
	"coderhub/shared/validator"
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
	if err := validator.New().Username(in.Username).Password(in.PasswordHash).Check(); err != nil {
		return nil, err
	}

	exists, _ := NewCheckUserExistsLogic(l.ctx, l.svcCtx).CheckUserExists(&user.CheckUserExistsRequest{Username: in.Username})
	if exists.Exists {
		return nil, errors.New("用户已存在")
	}

	ID := snowFlake.GenID()
	Password, _ := bcryptUtil.PasswordHash(in.PasswordHash)
	if err := l.svcCtx.UserRepository.CreateUser(&model.User{
		ID:       ID,
		UserName: in.Username,
		Password: Password,
	}); err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		UserId: ID,
	}, nil
}
