package logic

import (
	"coderhub/model"
	"coderhub/rpc/TechSphere/Articles/articles"
	"coderhub/rpc/TechSphere/Articles/internal/svc"
	"coderhub/shared/SnowFlake"
	"coderhub/shared/Validator"
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateArticleLogic {
	return &CreateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// validateCreateArticleRequest 验证文章创建请求
func (l *CreateArticleLogic) validateCreateArticleRequest(req *articles.CreateArticleRequest) error {
	// 验证基本字段
	if err := Validator.New().
		Title(req.Title).
		Summary(req.Summary).
		Content(req.Content).
		Tags(req.Tags).
		Check(); err != nil {
		return fmt.Errorf("字段验证失败: %w", err)
	}

	// 验证图片数量
	if len(req.ImageIds) > maxImageCount {
		return fmt.Errorf("图片数量不能超过%d张", maxImageCount)
	}

	return nil
}

// CreateArticle 创建文章
func (l *CreateArticleLogic) CreateArticle(in *articles.CreateArticleRequest) (*articles.CreateArticleResponse, error) {
	// 验证请求参数
	if err := l.validateCreateArticleRequest(in); err != nil {
		return nil, fmt.Errorf("请求参数验证失败: %w", err)
	}

	// 生成文章ID
	articleID := SnowFlake.GenID()

	// 创建文章模型
	article := &model.Articles{
		ID:       articleID,
		Title:    in.Title,
		Summary:  in.Summary,
		Content:  in.Content,
		Tags:     strings.Join(in.Tags, ","),
		Status:   in.Status,
		Type:     in.Type,
		AuthorID: in.AuthorId,
	}

	// 保存文章
	if err := l.svcCtx.ArticleRepository.CreateArticle(article); err != nil {
		return nil, fmt.Errorf("保存文章失败: %w", err)
	}

	return &articles.CreateArticleResponse{
		Id: articleID,
	}, nil
}
