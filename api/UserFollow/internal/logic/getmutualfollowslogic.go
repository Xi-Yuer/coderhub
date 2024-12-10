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

type GetMutualFollowsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取互相关注列表
func NewGetMutualFollowsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMutualFollowsLogic {
	return &GetMutualFollowsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMutualFollowsLogic) GetMutualFollows() (resp *types.GetMutualFollowsResp, err error) {
	UserID, err := utils.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err)
	}

	mutualFollowsResp, err := l.svcCtx.UserFollowService.GetMutualFollows(l.ctx, &userfollowservice.GetMutualFollowsReq{
		UserId: UserID,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp(mutualFollowsResp)
}

func (l *GetMutualFollowsLogic) successResp(mutualFollowsResp *userfollowservice.GetMutualFollowsResp) (*types.GetMutualFollowsResp, error) {
	userFollowList := make([]types.UserFollowInfo, 0, len(mutualFollowsResp.MutualFollows))
	for _, mutualFollow := range mutualFollowsResp.MutualFollows {
		userFollowList = append(userFollowList, types.UserFollowInfo{
			UserId:   mutualFollow.Id,
			Username: mutualFollow.Username,
			Avatar:   mutualFollow.Avatar,
		})
	}
	return &types.GetMutualFollowsResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: &types.UserFollowList{
			List:  userFollowList,
			Total: mutualFollowsResp.Total,
		},
	}, nil
}

func (l *GetMutualFollowsLogic) errorResp(err error) (*types.GetMutualFollowsResp, error) {
	return &types.GetMutualFollowsResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
