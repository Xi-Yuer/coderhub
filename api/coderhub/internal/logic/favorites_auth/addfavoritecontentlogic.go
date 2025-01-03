package favorites_auth

import (
	"coderhub/conf"
	"coderhub/rpc/coderhub/coderhub"
	"coderhub/shared/utils"
	"context"

	"coderhub/api/coderhub/internal/svc"
	"coderhub/api/coderhub/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFavoriteContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewAddFavoriteContentLogic 添加收藏内容
func NewAddFavoriteContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFavoriteContentLogic {
	return &AddFavoriteContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFavoriteContentLogic) AddFavoriteContent(req *types.CreateFavorReq) (resp *types.CreateFavorResp, err error) {
	userID, err := utils.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err)
	}
	response, err := l.svcCtx.FavoriteContentService.CreateFavor(l.ctx, &coderhub.CreateFavorRequest{
		UserId:        userID,
		FavorFolderId: utils.GenID(),
		EntityId:      req.EntityId,
		EntityType:    req.EntityType,
	})
	if err != nil {
		return l.errorResp(err)
	}
	return l.successResp(response)
}

func (l *AddFavoriteContentLogic) successResp(createFavorResp *coderhub.CreateFavorResponse) (*types.CreateFavorResp, error) {
	return &types.CreateFavorResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: createFavorResp.Success,
	}, nil
}

func (l *AddFavoriteContentLogic) errorResp(err error) (*types.CreateFavorResp, error) {
	return &types.CreateFavorResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
		Data: false,
	}, nil
}
