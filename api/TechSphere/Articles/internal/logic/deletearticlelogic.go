package logic

import (
	"coderhub/shared/utils"
	"context"
	"strconv"

	"coderhub/api/TechSphere/Articles/internal/svc"
	"coderhub/api/TechSphere/Articles/internal/types"
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Articles/articles"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteArticleLogic) DeleteArticle(req *types.DeleteArticleReq) (*types.DeleteArticleResp, error) {
	_, err := utils.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err), nil
	}
	ctx := utils.SetUserMetaData(l.ctx) // 设置元数据

	if err := utils.NewValidator().ArticleID(req.Id).Check(); err != nil {
		return l.errorResp(err), nil
	}

	articleIdInt, err := strconv.ParseInt(strconv.FormatInt(req.Id, 10), 10, 64)
	if err != nil {
		return l.errorResp(err), nil
	}

	if err := l.deleteArticle(ctx, articleIdInt); err != nil {
		return l.errorResp(err), nil
	}

	return l.successResp(), nil
}

func (l *DeleteArticleLogic) errorResp(err error) *types.DeleteArticleResp {
	return &types.DeleteArticleResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
		Data: false,
	}
}

func (l *DeleteArticleLogic) successResp() *types.DeleteArticleResp {
	return &types.DeleteArticleResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: true,
	}
}

func (l *DeleteArticleLogic) deleteArticle(ctx context.Context, articleId int64) error {
	userID, err := utils.GetUserID(l.ctx)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.ArticleService.DeleteArticle(ctx, &articles.DeleteArticleRequest{
		Id:     articleId,
		UserId: userID,
	})
	return err
}
