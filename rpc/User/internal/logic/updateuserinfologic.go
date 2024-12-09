package logic

import (
	"coderhub/model"
	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"
	"coderhub/shared/MetaData"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error) {
	if userId, err := MetaData.GetUserMetaData(l.ctx); err != nil || userId != strconv.FormatInt(in.UserId, 10) {
		return nil, errors.New("非法操作")
	}
	UserInfo, err := NewGetUserInfoLogic(l.ctx, l.svcCtx).GetUserInfo(&user.GetUserInfoRequest{
		UserId: in.UserId,
	})
	if err != nil {
		return nil, err
	}
	if UserInfo.UserId == 0 {
		return nil, errors.New("用户不存在")
	}

	fmt.Printf("in:%#v\n", in)
	oldUser, err := l.svcCtx.UserRepository.GetUserByID(UserInfo.UserId)
	if err != nil {
		return nil, err
	}
	if err := l.svcCtx.UserRepository.UpdateUser(&model.User{
		ID:       UserInfo.UserId,
		UserName: oldUser.UserName,
		Password: oldUser.Password,
		NickName: sql.NullString{String: in.Nickname, Valid: in.Nickname != ""},
		Email:    sql.NullString{String: in.Email, Valid: in.Email != ""},
		Avatar:   oldUser.Avatar,
		Status:   oldUser.Status,
		IsAdmin:  oldUser.IsAdmin,
	}); err != nil {
		return nil, err
	}

	return &user.UpdateUserInfoResponse{
		Success: true,
	}, nil
}
