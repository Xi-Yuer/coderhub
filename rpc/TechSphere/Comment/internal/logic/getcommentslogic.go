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

type GetCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetComments 获取评论列表
func (l *GetCommentsLogic) GetComments(in *comment.GetCommentsRequest) (*comment.GetCommentsResponse, error) {
	comments, total, err := l.svcCtx.CommentRepository.ListByArticleID(l.ctx, in.ArticleId, int64(in.Page), int64(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 构建树形结构
	rootComments := l.buildTree(comments)

	return &comment.GetCommentsResponse{
		Comments: rootComments,
		Total:    int32(total),
	}, nil
}

// buildTree 构建树形结构
func (l *GetCommentsLogic) buildTree(comments []model.Comment) []*comment.Comment {
	if len(comments) == 0 {
		return nil
	}

	// 收集所有评论ID
	commentIds := make([]int64, len(comments))
	for i, val := range comments {
		commentIds[i] = val.ID
	}

	// 批量获取所有评论的图片关联
	imageRelations, err := l.svcCtx.ImageRelationService.BatchGetImagesByEntity(l.ctx, &imageRelation.BatchGetImagesByEntityRequest{
		EntityIds:  commentIds,
		EntityType: model.ImageRelationComment,
	})
	if err != nil {
		// 添加错误日志
		l.Logger.Errorf("获取评论图片失败: %v", err)
	}

	// 构建评论ID到图片列表的映射
	commentImages := make(map[int64][]*comment.CommentImage)
	if err == nil {
		for _, img := range imageRelations.Relations {
			imageId := strconv.FormatInt(img.ImageId, 10)
			commentImages[img.EntityId] = append(commentImages[img.EntityId], &comment.CommentImage{
				ImageId:      imageId,
				Url:          img.Url,
				ThumbnailUrl: img.ThumbnailUrl,
			})
		}
	}

	rootComments := make([]*comment.Comment, len(comments))
	for i, val := range comments {
		rootComments[i] = &comment.Comment{
			Id:        val.ID,
			ArticleId: val.ArticleID,
			Content:   val.Content,
			ParentId:  val.ParentID,
			UserId:    val.UserID,
			Replies:   l.buildTree(val.Replies),
			LikeCount: val.LikeCount,
			Images:    commentImages[val.ID], // 从映射中获取图片
			CreatedAt: val.CreatedAt.Unix(),
			UpdatedAt: val.UpdatedAt.Unix(),
		}
	}
	return rootComments
}
