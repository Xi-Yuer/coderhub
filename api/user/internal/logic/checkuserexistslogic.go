package logic

import (
	"context"

	"coderhub/api/user/internal/svc"
	"coderhub/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserExistsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckUserExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserExistsLogic {
	return &CheckUserExistsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckUserExistsLogic) CheckUserExists(req *types.CheckUserExistsRequest) (resp *types.CheckUserExistsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
