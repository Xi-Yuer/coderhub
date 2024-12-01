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

	// 获取图片关联（文章配图）
	images, err := l.svcCtx.ImageRelationService.BatchGetImagesByEntity(l.ctx, &imageRelation.BatchGetImagesByEntityRequest{
		EntityIds:  []int64{article.ID},
		EntityType: model.ImageRelationArticleContent,
	})
	if err != nil {
		l.Logger.Errorf("获取图片关联失败: %v", err)
		return nil, fmt.Errorf("获取图片关联失败: %v", err)
	}
	l.Logger.Info("RPC: 获取文章配图成功, 配图数量:", len(images.Relations))
	// 获取图片关联（文章封面）
	coverImages, err := l.svcCtx.ImageRelationService.BatchGetImagesByEntity(l.ctx, &imageRelation.BatchGetImagesByEntityRequest{
		EntityIds:  []int64{article.ID},
		EntityType: model.ImageRelationArticleCover,
	})
	if err != nil {
		l.Logger.Errorf("获取图片关联失败: %v", err)
		return nil, fmt.Errorf("获取图片关联失败: %v", err)
	}
	l.Logger.Info("RPC: 获取文章封面成功, 封面数量:", len(coverImages.Relations))
	// 获取文章配图
	articleImages := make([]*articles.Image, 0)
	if len(images.Relations) > 0 {
		for _, image := range images.Relations {
			imageId := strconv.FormatInt(image.ImageId, 10)
			articleImages = append(articleImages, &articles.Image{
				ImageId:      imageId,
				Url:          image.Url,
				ThumbnailUrl: image.ThumbnailUrl,
			})
		}
	}
	var coverImage *articles.Image
	if len(coverImages.Relations) > 0 {
		coverImage = &articles.Image{
			ImageId:      strconv.FormatInt(coverImages.Relations[0].ImageId, 10),
			Url:          coverImages.Relations[0].Url,
			ThumbnailUrl: coverImages.Relations[0].ThumbnailUrl,
		}
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
