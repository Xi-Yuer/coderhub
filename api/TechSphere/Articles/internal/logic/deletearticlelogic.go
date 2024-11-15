package logic

import (
	"coderhub/shared/MetaData"
	"context"
	"strconv"

	"coderhub/api/TechSphere/Articles/internal/svc"
	"coderhub/api/TechSphere/Articles/internal/types"
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Articles/articles"
	"coderhub/shared/Validator"

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
	if err := Validator.New().ArticleID(req.Id).Check(); err != nil {
		return l.errorResp(err), nil
	}

	articleIdInt, err := strconv.ParseInt(strconv.FormatInt(req.Id, 10), 10, 64)
	if err != nil {
		return l.errorResp(err), nil
	}

	if err := l.deleteArticle(articleIdInt); err != nil {
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

func (l *DeleteArticleLogic) deleteArticle(articleId int64) error {
	userID, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return err
	}
	_, err = l.svcCtx.ArticleService.DeleteArticle(l.ctx, &articles.DeleteArticleRequest{
		Id:     articleId,
		UserId: userID,
	})
	return err
}
