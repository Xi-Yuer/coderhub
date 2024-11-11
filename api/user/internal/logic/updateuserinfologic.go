package logic

import (
	"coderhub/rpc/user/user"
	"coderhub/shared/metaData"
	"context"

	"coderhub/api/user/internal/svc"
	"coderhub/api/user/internal/types"

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
	ctx := metaData.SetUserMetaData(l.ctx)
	_, err = l.svcCtx.UserService.UpdateUserInfo(ctx, &user.UpdateUserInfoRequest{
		UserId:   req.UserId,
		Email:    req.Email,
		Nickname: req.Nickname,
	})
	return
}
