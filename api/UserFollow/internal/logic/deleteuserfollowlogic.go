package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"
	"coderhub/conf"
	"coderhub/rpc/UserFollow/userfollowservice"
	"coderhub/shared/MetaData"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消关注
func NewDeleteUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserFollowLogic {
	return &DeleteUserFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserFollowLogic) DeleteUserFollow(req *types.DeleteUserFollowReq) (resp *types.DeleteUserFollowResp, err error) {
	UserID, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err)
	}
	_, err = l.svcCtx.UserFollowService.DeleteUserFollow(l.ctx, &userfollowservice.DeleteUserFollowReq{
		FollowerId: UserID,
		FollowedId: req.FollowedId,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp()
}

func (l *DeleteUserFollowLogic) successResp() (*types.DeleteUserFollowResp, error) {
	return &types.DeleteUserFollowResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
	}, nil
}

func (l *DeleteUserFollowLogic) errorResp(err error) (*types.DeleteUserFollowResp, error) {
	return &types.DeleteUserFollowResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
