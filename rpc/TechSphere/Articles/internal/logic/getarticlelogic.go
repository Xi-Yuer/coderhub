package logic

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"coderhub/model"
	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/TechSphere/Articles/articles"
	"coderhub/rpc/TechSphere/Articles/internal/svc"
	"coderhub/shared/Validator"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *GetArticleLogic) GetArticle(in *articles.GetArticleRequest) (*articles.GetArticleResponse, error) {
	// 参数校验
	if err := Validator.New().ArticleID(in.Id).Check(); err != nil {
		return nil, errors.New(err.Error())
	}

	// 从数据库获取文章
	article, err := l.svcCtx.ArticleRepository.GetArticleByID(in.Id)
	if err != nil {
		l.Logger.Errorf("获取文章失败: %v", err)
		return nil, fmt.Errorf("获取文章失败: %v", err)
	}

	if article == nil {
		return nil, fmt.Errorf("文章不存在")
	}

	// 获取图片关联
	images, err := l.svcCtx.ImageRelationService.GetImagesByEntity(l.ctx, &imageRelation.GetImagesByEntityRequest{
		EntityId:   article.ID,
		EntityType: model.ImageRelationArticleContent,
	})
	if err != nil {
		l.Logger.Errorf("获取图片关联失败: %v", err)
		return nil, fmt.Errorf("获取图片关联失败: %v", err)
	}
	// 获取封面图片
	var coverImage *articles.Image
	if len(images.Images) > 0 {
		imageId := strconv.FormatInt(images.Images[0].ImageId, 10)
		coverImage = &articles.Image{
			ImageId:      imageId,
			Url:          images.Images[0].Url,
			ThumbnailUrl: images.Images[0].ThumbnailUrl,
			Width:        images.Images[0].Width,
			Height:       images.Images[0].Height,
		}
	}
	// 获取文章配图
	articleImages := make([]*articles.Image, 0)
	for _, image := range images.Images[1:] {
		imageId := strconv.FormatInt(image.ImageId, 10)
		articleImages = append(articleImages, &articles.Image{
			ImageId:      imageId,
			Url:          image.Url,
			ThumbnailUrl: image.ThumbnailUrl,
			Width:        image.Width,
			Height:       image.Height,
		})
	}

	// 转换为响应格式
	response := &articles.GetArticleResponse{
		Article: &articles.Article{
			Id:           article.ID,
			Type:         article.Type,
			Title:        article.Title,
			Content:      article.Content,
			Summary:      article.Summary,
			Images:       articleImages,
			CoverImage:   coverImage,
			AuthorId:     article.AuthorID,
			Tags:         strings.Split(article.Tags, ","),
			ViewCount:    article.ViewCount,
			LikeCount:    article.LikeCount,
			CommentCount: article.CommentCount,
			Status:       article.Status,
			CreatedAt:    article.CreatedAt.Unix(),
			UpdatedAt:    article.UpdatedAt.Unix(),
		},
	}

	return response, nil
}
