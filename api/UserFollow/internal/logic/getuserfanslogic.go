package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
