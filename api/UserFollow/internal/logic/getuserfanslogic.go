package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"
	"coderhub/conf"
	"coderhub/rpc/UserFollow/userfollowservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFansLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户粉丝列表
func NewGetUserFansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFansLogic {
	return &GetUserFansLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserFansLogic) GetUserFans(req *types.GetUserFansReq) (resp *types.GetUserFansResp, err error) {
	userFansResp, err := l.svcCtx.UserFollowService.GetUserFans(l.ctx, &userfollowservice.GetUserFansReq{
		FollowedId: req.UserId,
		Page:       req.Page,
		PageSize:   req.PageSize,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp(userFansResp)
}

func (l *GetUserFansLogic) successResp(userFansResp *userfollowservice.GetUserFansResp) (*types.GetUserFansResp, error) {
	userFansList := make([]types.UserFollowInfo, 0, len(userFansResp.UserFans))
	for _, userFan := range userFansResp.UserFans {
		userFansList = append(userFansList, types.UserFollowInfo{
			UserId:   userFan.Id,
			Username: userFan.Username,
			Avatar:   userFan.Avatar,
		})
	}

	return &types.GetUserFansResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: &types.UserFollowedList{
			List:  userFansList,
			Total: userFansResp.Total,
		},
	}, nil
}

func (l *GetUserFansLogic) errorResp(err error) (*types.GetUserFansResp, error) {
	return &types.GetUserFansResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
		Data: &types.UserFollowedList{
			List:  nil,
			Total: 0,
		},
	}, nil
}
