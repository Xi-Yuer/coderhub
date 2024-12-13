package academic_public

import (
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAcademicNavigatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取学术导航
func NewGetAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAcademicNavigatorLogic {
	return &GetAcademicNavigatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAcademicNavigatorLogic) GetAcademicNavigator(req *types.GetAcademicNavigatorReq) (resp *types.GetAcademicNavigatorResp, err error) {
	// todo: add your logic here and delete this line

	return
}
