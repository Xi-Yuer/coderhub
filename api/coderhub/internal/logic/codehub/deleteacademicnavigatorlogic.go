package codehub

import (
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAcademicNavigatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除学术导航
func NewDeleteAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAcademicNavigatorLogic {
	return &DeleteAcademicNavigatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAcademicNavigatorLogic) DeleteAcademicNavigator(req *types.DeleteAcademicNavigatorReq) (resp *types.DeleteAcademicNavigatorResp, err error) {
	// todo: add your logic here and delete this line

	return
}
