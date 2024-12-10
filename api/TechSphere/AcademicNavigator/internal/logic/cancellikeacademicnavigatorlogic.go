package logic

import (
	"context"

	"coderhub/api/TechSphere/AcademicNavigator/internal/svc"
	"coderhub/api/TechSphere/AcademicNavigator/internal/types"
	"coderhub/conf"
	"coderhub/rpc/TechSphere/AcademicNavigator/academic_navigator"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeAcademicNavigatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 取消点赞学术导航
func NewCancelLikeAcademicNavigatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeAcademicNavigatorLogic {
	return &CancelLikeAcademicNavigatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLikeAcademicNavigatorLogic) CancelLikeAcademicNavigator(req *types.CancelLikeAcademicNavigatorReq) (resp *types.CancelLikeAcademicNavigatorResp, err error) {
	_, err = l.svcCtx.AcademicNavigatorService.CancelLikeAcademicNavigator(l.ctx, &academic_navigator.CancelLikeAcademicNavigatorRequest{
		Id:     req.Id,
		UserId: req.UserId,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp()
}

func (l *CancelLikeAcademicNavigatorLogic) successResp() (*types.CancelLikeAcademicNavigatorResp, error) {
	return &types.CancelLikeAcademicNavigatorResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: true,
	}, nil
}

func (l *CancelLikeAcademicNavigatorLogic) errorResp(err error) (*types.CancelLikeAcademicNavigatorResp, error) {
	return &types.CancelLikeAcademicNavigatorResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
		Data: false,
	}, nil
}
