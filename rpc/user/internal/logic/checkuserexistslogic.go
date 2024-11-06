package logic

import (
	"coderhub/model"
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

func (l *CheckUserExistsLogic) CheckUserExists(in *user.CheckUserExistsRequest) (*user.CheckUserExistsResponse, error) {
	if tx := l.svcCtx.SqlDB.First(&model.User{}, "user_name = ?", in.Username); tx.RowsAffected == 0 {
		return &user.CheckUserExistsResponse{
			Exists: false,
		}, nil
	}
	return &user.CheckUserExistsResponse{
		Exists: true,
	}, nil
}
