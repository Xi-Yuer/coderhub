package logic

import (
	"coderhub/model"
	"coderhub/shared/MetaData"
	"coderhub/shared/Validator"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"coderhub/rpc/TechSphere/Articles/articles"
	"coderhub/rpc/TechSphere/Articles/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteArticleLogic 处理删除文章的业务逻辑
type DeleteArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewDeleteArticleLogic 创建DeleteArticleLogic实例
func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteArticle 处理删除文章的请求
// 包含参数验证、文章存在性检查和软删除操作
func (l *DeleteArticleLogic) DeleteArticle(in *articles.DeleteArticleRequest) (*articles.DeleteArticleResponse, error) {
	// 使用 defer 记录方法执行时间
	defer func(start time.Time) {
		l.Logger.Infof("DeleteArticle 执行耗时: %v", time.Since(start))
	}(time.Now())

	// 从 metadata 中获取 userId
	var (
		userId string
		err    error
	)
	if userId, err = MetaData.GetUserMetaData(l.ctx); err != nil {
		return nil, err
	}

	if userId != strconv.FormatInt(in.UserId, 10) {
		return nil, fmt.Errorf("非法操作")
	}

	// 参数校验
	if err := l.validateRequest(in); err != nil {
		return nil, err
	}

	// 检查文章是否存在
	article, err := l.getArticle(in.Id)
	if err != nil {
		return nil, err
	}

	// 记录操作日志
	l.Logger.Infof("开始删除文章, ID: %d, 标题: %s", article.ID, article.Title)

	// 权限校验
	if article.AuthorID != in.UserId {
		l.Logger.Errorf(ErrValidationFailed)
		return nil, errors.New(ErrValidationFailed)
	}

	// 执行软删除
	if err := l.deleteArticle(in.Id); err != nil {
		return nil, err
	}

	return &articles.DeleteArticleResponse{Success: true}, nil
}

// validateRequest 验证删除请求的参数
func (l *DeleteArticleLogic) validateRequest(in *articles.DeleteArticleRequest) error {
	if err := Validator.New().ArticleID(in.Id).Check(); err != nil {
		l.Logger.Errorf(ErrValidationFailed, err)
		return fmt.Errorf(ErrValidationFailed, err)
	}
	return nil
}

// getArticle 获取并验证文章是否存在
func (l *DeleteArticleLogic) getArticle(articleID int64) (*model.Articles, error) {
	article, err := l.svcCtx.ArticleRepository.GetArticleByID(articleID)
	if err != nil {
		l.Logger.Errorf(ErrGetArticleFailed, err)
		return nil, fmt.Errorf(ErrGetArticleFailed, err)
	}

	if article == nil {
		l.Logger.Errorf(ErrArticleNotFound)
		return nil, errors.New(ErrArticleNotFound)
	}

	return article, nil
}

// deleteArticle 执行文章删除操作
func (l *DeleteArticleLogic) deleteArticle(articleID int64) error {
	if err := l.svcCtx.ArticleRepository.DeleteArticle(articleID); err != nil {
		l.Logger.Errorf(ErrDeleteArticleFailed, err)
		return fmt.Errorf(ErrDeleteArticleFailed, err)
	}
	return nil
}
