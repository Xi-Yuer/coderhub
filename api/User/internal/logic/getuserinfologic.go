package logic

import (
	"coderhub/conf"
	"coderhub/rpc/User/user"
	"context"

	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	UserInfo, err := l.svcCtx.UserService.GetUserInfo(l.ctx, &user.GetUserInfoRequest{UserId: req.UserId})
	if err != nil {
		return &types.GetUserInfoResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
			Data: nil,
		}, nil
	}
	return &types.GetUserInfoResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: &types.UserInfo{
			UserId:    UserInfo.UserId,
			Username:  UserInfo.UserName,
			Avatar:    UserInfo.Avatar,
			Email:     UserInfo.Email,
			Nickname:  UserInfo.NickName,
			IsAdmin:   UserInfo.IsAdmin,
			Status:    UserInfo.Status,
			CreatedAt: UserInfo.CreatedAt,
			UpdatedAt: UserInfo.UpdatedAt,
		},
	}, nil
}
