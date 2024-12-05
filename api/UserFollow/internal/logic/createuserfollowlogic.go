package logic

import (
	"context"

	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
