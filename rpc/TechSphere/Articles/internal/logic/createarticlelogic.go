package logic

import (
	"coderhub/model"
	"coderhub/rpc/TechSphere/Articles/articles"
	"coderhub/rpc/TechSphere/Articles/internal/svc"
	"coderhub/shared/SnowFlake"
	"coderhub/shared/Validator"
	"context"
	"encoding/json"
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

// validateImageURL 验证图片URL是否合法
func validateImageURL(url string) bool {
	return strings.HasPrefix(url, urlHTTP) || strings.HasPrefix(url, urlHTTPS)
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

	// 验证图片URL
	if len(req.ImageUrls) > maxImageCount {
		return fmt.Errorf("图片数量不能超过%d张", maxImageCount)
	}

	for _, url := range req.ImageUrls {
		if !validateImageURL(url) {
			return fmt.Errorf("图片URL格式不正确: %s", url)
		}
	}

	// 验证封面图
	if req.CoverImage != "" && !validateImageURL(req.CoverImage) {
		return fmt.Errorf("封面图URL格式不正确: %s", req.CoverImage)
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
	imageURLsJSON, err := json.Marshal(in.ImageUrls)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ImageUrls: %v", err)
	}
	article := &model.Articles{
		ID:         articleID,
		Title:      in.Title,
		Summary:    in.Summary,
		Content:    in.Content,
		Tags:       strings.Join(in.Tags, ","),
		Status:     in.Status,
		Type:       in.Type,
		AuthorID:   in.AuthorId,
		ImageURLs:  string(imageURLsJSON),
		CoverImage: in.CoverImage,
	}

	// 保存文章
	if err := l.svcCtx.ArticleRepository.CreateArticle(article); err != nil {
		return nil, fmt.Errorf("保存文章失败: %w", err)
	}

	return &articles.CreateArticleResponse{
		Id: articleID,
	}, nil
}
