package logic

import (
	"context"
	"strconv"

	"coderhub/model"
	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/TechSphere/Comment/comment"
	"coderhub/rpc/TechSphere/Comment/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetComment 获取单个评论详情
func (l *GetCommentLogic) GetComment(in *comment.GetCommentRequest) (*comment.GetCommentResponse, error) {
	commentModel, err := l.svcCtx.CommentRepository.GetByID(l.ctx, in.CommentId)
	if err != nil {
		return nil, err
	}
	// 获取图片关联
	imageRelations, err := l.svcCtx.ImageRelationService.GetImagesByEntity(l.ctx, &imageRelation.GetImagesByEntityRequest{
		EntityId:   in.CommentId,
		EntityType: model.ImageRelation_COMMENT,
	})
	if err != nil {
		return nil, err
	}
	images := make([]*comment.CommentImage, 0)
	for _, imageRelation := range imageRelations.Images {
		imageId := strconv.FormatInt(imageRelation.ImageId, 10)
		images = append(images, &comment.CommentImage{
			ImageId:      imageId,
			Url:          imageRelation.Url,
			ThumbnailUrl: imageRelation.ThumbnailUrl,
		})
	}
	return &comment.GetCommentResponse{
		Comment: &comment.Comment{
			Id:        commentModel.ID,
			ArticleId: commentModel.ArticleID,
			Content:   commentModel.Content,
			ParentId:  commentModel.ParentID,
			UserId:    commentModel.UserID,
			CreatedAt: commentModel.CreatedAt.Unix(),
			UpdatedAt: commentModel.UpdatedAt.Unix(),
			LikeCount: commentModel.LikeCount,
			Images:    images,
		},
	}, nil
}
