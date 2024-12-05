package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"
	"coderhub/conf"
	"coderhub/rpc/UserFollow/userfollowservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户关注列表
func NewGetUserFollowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowsLogic {
	return &GetUserFollowsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserFollowsLogic) GetUserFollows(req *types.GetUserFollowsReq) (resp *types.GetUserFollowsResp, err error) {
	userFollowsResp, err := l.svcCtx.UserFollowService.GetUserFollows(l.ctx, &userfollowservice.GetUserFollowsReq{
		FollowerId: req.UserId,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return l.errorResp(err)
	}
	return l.successResp(userFollowsResp)
}

func (l *GetUserFollowsLogic) successResp(userFollowsResp *userfollowservice.GetUserFollowsResp) (*types.GetUserFollowsResp, error) {
	userFollowsList := make([]types.UserFollowInfo, 0, len(userFollowsResp.UserFollows))
	for _, userFollow := range userFollowsResp.UserFollows {
		userFollowsList = append(userFollowsList, types.UserFollowInfo{
			UserId:   userFollow.Id,
			Username: userFollow.Username,
			Avatar:   userFollow.Avatar,
		})
	}
	return &types.GetUserFollowsResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: &types.UserFollowList{
			List:  userFollowsList,
			Total: userFollowsResp.Total,
		},
	}, nil
}

func (l *GetUserFollowsLogic) errorResp(err error) (*types.GetUserFollowsResp, error) {
	return &types.GetUserFollowsResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
