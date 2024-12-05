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

type CreateUserFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建关注关系
func NewCreateUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserFollowLogic {
	return &CreateUserFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserFollowLogic) CreateUserFollow(req *types.CreateUserFollowReq) (resp *types.CreateUserFollowResp, err error) {
	UserID, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err)
	}
	_, err = l.svcCtx.UserFollowService.CreateUserFollow(l.ctx, &userfollowservice.CreateUserFollowReq{
		FollowerId: UserID,
		FollowedId: req.FollowedId,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp()
}

func (l *CreateUserFollowLogic) successResp() (*types.CreateUserFollowResp, error) {
	return &types.CreateUserFollowResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: true,
	}, nil
}

func (l *CreateUserFollowLogic) errorResp(err error) (*types.CreateUserFollowResp, error) {
	return &types.CreateUserFollowResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
