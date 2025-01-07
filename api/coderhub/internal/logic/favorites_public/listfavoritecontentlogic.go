package favorites_public

import (
	"coderhub/conf"
	"coderhub/rpc/coderhub/coderhub"
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFavoriteContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewListFavoriteContentLogic 获取收藏内容列表
func NewListFavoriteContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFavoriteContentLogic {
	return &ListFavoriteContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFavoriteContentLogic) ListFavoriteContent(req *types.GetFavorListReq) (resp *types.GetFavorListResp, err error) {
	if err != nil {
		return l.errorResp(err)
	}
	list, err := l.svcCtx.FavoriteContentService.GetFavorList(l.ctx, &coderhub.GetFavorListRequest{
		Page:          req.Page,
		PageSize:      req.PageSize,
		UserId:        req.UserId,
		EntityType:    req.EntityType,
		FavorFolderId: req.FavorFoldId,
	})
	if err != nil {
		return l.errorResp(err)
	}

	response := make([]*types.Favor, 0, len(list.Favors))

	for _, v := range list.Favors {
		response = append(response, &types.Favor{
			ID:          v.Id,
			CreateUser:  v.UserId,
			FavorFoldId: v.FavorFolderId,
			EntityId:    v.EntityId,
			EntityType:  v.EntityType,
			CreatedAt:   v.CreateTime,
		})
	}

	return l.successResp(types.FavorList{
		Total: list.Total,
		List:  response,
	})
}

func (l *ListFavoriteContentLogic) successResp(list types.FavorList) (*types.GetFavorListResp, error) {
	return &types.GetFavorListResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: list,
	}, nil
}

func (l *ListFavoriteContentLogic) errorResp(err error) (*types.GetFavorListResp, error) {
	return &types.GetFavorListResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
		Data: types.FavorList{},
	}, nil
}
