package logic

import (
	"context"

	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetUserByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetUserByIDLogic {
	return &BatchGetUserByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量获取用户信息
func (l *BatchGetUserByIDLogic) BatchGetUserByID(in *user.BatchGetUserByIDRequest) (*user.BatchGetUserByIDResponse, error) {
	// todo: add your logic here and delete this line

	return &user.BatchGetUserByIDResponse{}, nil
}
