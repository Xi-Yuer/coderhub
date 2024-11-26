package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"coderhub/model"
	"coderhub/rpc/TechSphere/Articles/articles"
	"coderhub/rpc/TechSphere/Articles/internal/svc"
	"coderhub/shared/MetaData"
	"coderhub/shared/Validator"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	maxImageCount = 20
	urlHTTP       = "http://"
	urlHTTPS      = "https://"
)

// 错误常量
const (
	ErrValidationFailed    = "参数校验失败: %v"
	ErrUserMetaFailed      = "获取用户元数据失败: %v"
	ErrUserIDConversion    = "转换用户ID失败: %v"
	ErrArticleNotFound     = "文章不存在"
	ErrNoPermission        = "您无权修改此文章"
	ErrUpdateFailed        = "更新文章失败: %v"
	ErrImageCountExceeded  = "图片数量不能超过%d张"
	ErrInvalidImageURL     = "图片URL格式不正确"
	ErrInvalidCoverURL     = "封面图URL格式不正确"
	ErrGetArticleFailed    = "获取文章失败: %v"
	ErrDeleteArticleFailed = "删除文章失败: %v"
)

// UpdateArticleLogic 处理更新文章的业务逻辑
type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateArticle 更新文章
func (l *UpdateArticleLogic) UpdateArticle(in *articles.UpdateArticleRequest) (*articles.UpdateArticleResponse, error) {
	// 记录执行时间
	defer func(start time.Time) {
		l.Logger.Infof("UpdateArticle 执行耗时: %v", time.Since(start))
	}(time.Now())

	// 参数校验
	if err := l.validateArticleUpdate(in); err != nil {
		l.Logger.Errorf(ErrValidationFailed, err)
		return nil, err
	}

	// 获取用户ID
	userID, err := l.getUserID()
	if err != nil {
		return nil, err
	}

	// 检查文章权限
	article, err := l.checkArticlePermission(in.Id, userID)
	if err != nil {
		return nil, err
	}

	// 更新文章信息
	l.updateArticleFields(article, in)

	// 保存更新
	if err := l.svcCtx.ArticleRepository.UpdateArticle(article); err != nil {
		l.Logger.Errorf(ErrUpdateFailed, err)
		return nil, fmt.Errorf(ErrUpdateFailed, err)
	}

	l.Logger.Infof("文章更新成功, ID: %d", in.Id)
	return &articles.UpdateArticleResponse{Success: true}, nil
}

// validateArticleUpdate 验证文章更新请求
func (l *UpdateArticleLogic) validateArticleUpdate(req *articles.UpdateArticleRequest) error {
	// 基础字段验证
	if err := Validator.New().
		ArticleID(req.Id).
		Title(req.Title).
		Summary(req.Summary).
		Content(req.Content).
		Tags(req.Tags).
		Check(); err != nil {
		return fmt.Errorf(ErrValidationFailed, err)
	}

	// 验证图片数量
	if len(req.ImageIds) > maxImageCount {
		return fmt.Errorf(ErrImageCountExceeded, maxImageCount)
	}

	return nil
}

// getUserID 获取并转换用户ID
func (l *UpdateArticleLogic) getUserID() (int64, error) {
	userId, err := MetaData.GetUserMetaData(l.ctx)
	if err != nil {
		l.Logger.Errorf(ErrUserMetaFailed, err)
		return 0, fmt.Errorf(ErrUserMetaFailed, err)
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		l.Logger.Errorf(ErrUserIDConversion, err)
		return 0, fmt.Errorf(ErrUserIDConversion, err)
	}

	return userIdInt, nil
}

// checkArticlePermission 检查文章权限
func (l *UpdateArticleLogic) checkArticlePermission(articleID int64, userID int64) (*model.Articles, error) {
	article, err := l.svcCtx.ArticleRepository.GetArticleByID(articleID)
	if err != nil {
		l.Logger.Errorf("获取文章失败: %v", err)
		return nil, err
	}

	if article == nil {
		return nil, errors.New(ErrArticleNotFound)
	}

	if userID != article.AuthorID {
		return nil, errors.New(ErrNoPermission)
	}

	return article, nil
}

// updateArticleFields 更新文章字段
func (l *UpdateArticleLogic) updateArticleFields(article *model.Articles, in *articles.UpdateArticleRequest) {
	article.Title = in.Title
	article.Content = in.Content
	article.Summary = in.Summary
	article.Tags = strings.Join(in.Tags, ",")
	article.Status = string(in.Status)
}
