package academic_auth

import (
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAcademicNavigatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增学术导航
func NewAddAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAcademicNavigatorLogic {
	return &AddAcademicNavigatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddAcademicNavigatorLogic) AddAcademicNavigator(req *types.AddAcademicNavigatorReq) (resp *types.AddAcademicNavigatorResp, err error) {
	// todo: add your logic here and delete this line

	return
}
