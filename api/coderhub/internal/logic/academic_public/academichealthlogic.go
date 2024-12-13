package academic_public

import (
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AcademicHealthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 健康检查
func NewAcademicHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AcademicHealthLogic {
	return &AcademicHealthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AcademicHealthLogic) AcademicHealth() (resp *types.HealthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
