package logic

import (
	"coderhub/conf"
	"coderhub/rpc/User/user"
	"coderhub/shared/MetaData"
	"context"

	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoRequest) (resp *types.UpdateUserInfoResponse, err error) {
	ctx := MetaData.SetUserMetaData(l.ctx)
	userID, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return &types.UpdateUserInfoResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
		}, nil
	}

	if _, err = l.svcCtx.UserService.UpdateUserInfo(ctx, &user.UpdateUserInfoRequest{
		UserId:   userID,
		Email:    req.Email,
		Nickname: req.Nickname,
	}); err != nil {
		return &types.UpdateUserInfoResponse{
			Response: types.Response{
				Code:    conf.HttpCode.HttpBadRequest,
				Message: err.Error(),
			},
		}, nil
	}

	return &types.UpdateUserInfoResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: true,
	}, nil
}
