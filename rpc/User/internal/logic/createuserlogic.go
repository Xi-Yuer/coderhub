package logic

import (
	"coderhub/model"
	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"
	"coderhub/shared/security"
	"coderhub/shared/utils"
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
	if err := utils.NewValidator().Username(in.Username).Password(in.PasswordHash).Check(); err != nil {
		return nil, err
	}

	exists, _ := NewCheckUserExistsLogic(l.ctx, l.svcCtx).CheckUserExists(&user.CheckUserExistsRequest{Username: in.Username})
	if exists.Exists {
		return nil, errors.New("用户已存在")
	}

	ID := utils.GenID()
	Password, _ := security.PasswordHash(in.PasswordHash)
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
