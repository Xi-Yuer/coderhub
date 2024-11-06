package logic

import (
	"coderhub/model"
	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"
	"coderhub/shared/bcryptUtil"
	"coderhub/shared/snowflake"
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
	checkUserExistsLogic := CheckUserExistsLogic{
		ctx:    l.ctx,
		svcCtx: l.svcCtx,
		Logger: l.Logger,
	}
	checkUserExists, _ := checkUserExistsLogic.CheckUserExists(&user.CheckUserExistsRequest{
		Username: in.Username,
	})
	if checkUserExists.Exists {
		return nil, errors.New("用户已存在")
	}

	ID := snowflake.GenID()
	// 密码加密
	Password, _ := bcryptUtil.PasswordHash(in.PasswordHash)
	if tx := l.svcCtx.SqlDB.Create(&model.User{
		ID:       ID,
		UserName: in.Username,
		Password: Password,
		NickName: nil,
		Email:    nil,
		Avatar:   nil,
		Status:   false,
		IsAdmin:  false,
	}); tx.Error != nil {
		return nil, tx.Error
	}

	return &user.CreateUserResponse{
		UserId: ID,
	}, nil
}
