package logic

import (
	"context"

	"coderhub/api/TechSphere/Articles/internal/svc"
	"coderhub/api/TechSphere/Articles/internal/types"
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Articles/articles"
	"coderhub/shared/MetaData"
	"coderhub/shared/Validator"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLikeCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLikeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLikeCountLogic {
	return &UpdateLikeCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLikeCountLogic) UpdateLikeCount(req *types.UpdateLikeCountReq) (resp *types.UpdateLikeCountResp, err error) {
	userId, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err), nil
	}
	ctx := MetaData.SetUserMetaData(l.ctx)

	if err := Validator.New().ArticleID(req.Id).Check(); err != nil {
		return l.errorResp(err), nil
	}

	if _, err := l.svcCtx.ArticleService.UpdateLikeCount(ctx, &articles.UpdateLikeCountRequest{
		Id:     req.Id,
		UserId: userId,
	}); err != nil {
		return l.errorResp(err), nil
	}

	return l.successResp(), nil
}

func (l *UpdateLikeCountLogic) errorResp(err error) *types.UpdateLikeCountResp {
	return &types.UpdateLikeCountResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
		Data: false,
	}
}

func (l *UpdateLikeCountLogic) successResp() *types.UpdateLikeCountResp {
	return &types.UpdateLikeCountResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: "success",
		},
		Data: true,
	}
}
