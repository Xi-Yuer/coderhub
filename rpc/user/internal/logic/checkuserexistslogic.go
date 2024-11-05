package logic

import (
	"context"

	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserExistsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserExistsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserExistsLogic {
	return &CheckUserExistsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckUserExists 检查用户是否存在
func (l *CheckUserExistsLogic) CheckUserExists(in *user.CheckUserExistsRequest) (*user.CheckUserExistsResponse, error) {
	return &user.CheckUserExistsResponse{
		Exists: true,
	}, nil
}
