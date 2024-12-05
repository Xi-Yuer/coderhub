package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"

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

func (l *GetMutualFollowsLogic) GetMutualFollows(req *types.GetMutualFollowsReq) (resp *types.GetMutualFollowsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
