package logic

import (
	"coderhub/shared/utils"
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"
	"coderhub/conf"
	"coderhub/rpc/UserFollow/userfollowservice"
	"github.com/zeromicro/go-zero/core/logx"
)

type IsUserFollowedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 检查是否关注
func NewIsUserFollowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsUserFollowedLogic {
	return &IsUserFollowedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsUserFollowedLogic) IsUserFollowed(req *types.IsUserFollowedReq) (resp *types.IsUserFollowedResp, err error) {
	UserID, err := utils.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err)
	}
	isUserFollowedResp, err := l.svcCtx.UserFollowService.IsUserFollowed(l.ctx, &userfollowservice.IsUserFollowedReq{
		FollowerId: UserID,
		FollowedId: req.FollowedId,
	})
	if err != nil {
		return l.errorResp(err)
	}
	return l.successResp(isUserFollowedResp)
}

func (l *IsUserFollowedLogic) successResp(isUserFollowedResp *userfollowservice.IsUserFollowedResp) (*types.IsUserFollowedResp, error) {
	return &types.IsUserFollowedResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: isUserFollowedResp.IsFollowed,
	}, nil
}

func (l *IsUserFollowedLogic) errorResp(err error) (*types.IsUserFollowedResp, error) {
	return &types.IsUserFollowedResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
