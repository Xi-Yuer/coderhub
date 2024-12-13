package user_public

import (
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordByLinkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过链接重置密码
func NewResetPasswordByLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordByLinkLogic {
	return &ResetPasswordByLinkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordByLinkLogic) ResetPasswordByLink(req *types.ResetPasswordByLinkReq) (resp *types.ResetPasswordByLinkResp, err error) {
	// todo: add your logic here and delete this line

	return
}
