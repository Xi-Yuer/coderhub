package logic

import (
	"coderhub/model"
	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"
	"coderhub/shared/metaData"
	"context"
	"database/sql"
	"errors"
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
	if userId, err := metaData.GetUserMetaData(l.ctx); err != nil || userId != strconv.FormatInt(in.UserId, 10) {
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

	if err := l.svcCtx.UserRepository.UpdateUser(&model.User{
		NickName: sql.NullString{
			String: in.Nickname,
		},
		Email: sql.NullString{
			String: in.Email,
		},
	}); err != nil {
		return nil, err
	}

	return &user.UpdateUserInfoResponse{
		Success: true,
	}, nil
}
