package logic

import (
	"coderhub/api/TechSphere/Articles/internal/svc"
	"coderhub/api/TechSphere/Articles/internal/types"
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Articles/articles"
	"coderhub/shared/MetaData"
	"coderhub/shared/Validator"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateArticleLogic) CreateArticle(req *types.CreateArticleReq) (*types.CreateArticleResp, error) {
	// 1. 参数验证
	if err := l.validateArticleData(req); err != nil {
		return l.errorResp(err), nil
	}

	// 2. 获取用户ID
	userId, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err), nil
	}

	// 3. 预处理文章内容
	article := l.prepareArticleData(req, userId)

	// 4. 创建文章
	articleId, err := l.createArticle(article)
	if err != nil {
		return l.errorResp(err), nil
	}

	// 5. 处理文章标签（异步）
	go l.handleArticleTags(articleId, req.Tags)

	return l.successResp(articleId), nil
}

func (l *CreateArticleLogic) errorResp(err error) *types.CreateArticleResp {
	return &types.CreateArticleResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}
}

func (l *CreateArticleLogic) successResp(articleId int64) *types.CreateArticleResp {
	return &types.CreateArticleResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: articleId,
	}
}

func (l *CreateArticleLogic) validateArticleData(req *types.CreateArticleReq) error {
	return Validator.New().
		Title(req.Title).
		Content(req.Content).
		ArticleType(req.Type).
		Tags(req.Tags).
		Check()
}

func (l *CreateArticleLogic) prepareArticleData(req *types.CreateArticleReq, userId int64) *articles.CreateArticleRequest {

	return &articles.CreateArticleRequest{
		Type:         req.Type,
		Title:        req.Title,
		Content:      req.Content,
		Summary:      l.generateSummary(req.Summary, req.Content),
		ImageIds:     nil,
		CoverImageId: "",
		AuthorId:     userId,
		Tags:         req.Tags,
		Status:       req.Status,
	}
}

func (l *CreateArticleLogic) generateSummary(summary, content string) string {
	if summary == "" {
		// TODO: 实现文章摘要生成逻辑
		// 可以截取前N个字符或使用AI生成摘要
		const summaryLength = 200
		if len(content) > summaryLength {
			return content[:summaryLength] + "..."
		}
		return content
	} else {
		return summary
	}
}

func (l *CreateArticleLogic) createArticle(article *articles.CreateArticleRequest) (int64, error) {
	resp, err := l.svcCtx.ArticleService.CreateArticle(l.ctx, article)
	if err != nil {
		return 0, err
	}
	return resp.Id, nil
}

func (l *CreateArticleLogic) handleArticleTags(articleId int64, tags []string) {
	// TODO: 实现标签处理逻辑
	// 1. 创建新标签（如果不存在）
	// 2. 建立文章和标签的关联关系
	// 3. 更新标签使用计数
}
