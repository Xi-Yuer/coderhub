package logic

import (
	"coderhub/model"
	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/TechSphere/Comment/comment"
	"coderhub/rpc/TechSphere/Comment/internal/svc"
	"coderhub/shared/SnowFlake"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateComment 创建评论
func (l *CreateCommentLogic) CreateComment(in *comment.CreateCommentRequest) (*comment.CreateCommentResponse, error) {
	CommentID := SnowFlake.GenID()
	commentModel := &model.Comment{
		ID:         CommentID,
		ArticleID:  in.ArticleId,
		Content:    in.Content,
		ParentID:   in.ParentId,
		UserID:     in.UserId,
		ReplyToUID: in.ReplyToUid,
	}
	// 如果评论有携带图片，则需要创建图片关联
	var imageRelationModels []*imageRelation.CreateRelationRequest
	if len(in.ImageIds) > 0 {
		for _, imageId := range in.ImageIds {
			imageIdInt, err := strconv.ParseInt(imageId, 10, 64)
			if err != nil {
				return nil, err
			}
			imageRelationModels = append(imageRelationModels, &imageRelation.CreateRelationRequest{
				ImageId:    imageIdInt,
				EntityId:   CommentID,
				EntityType: model.ImageRelation_COMMENT,
			})
		}
	}
	_, err := l.svcCtx.ImageRelationService.BatchCreateRelation(l.ctx, &imageRelation.BatchCreateRelationRequest{
		Relations: imageRelationModels,
	})
	if err != nil {
		return nil, err
	}
	// 创建评论
	if err := l.svcCtx.CommentRepository.Create(l.ctx, commentModel); err != nil {
		// 事务回滚
		l.svcCtx.ImageRelationService.BatchDeleteRelation(l.ctx, &imageRelation.BatchDeleteRelationRequest{
			Ids: []int64{CommentID},
		})
		return nil, err
	}
	return &comment.CreateCommentResponse{
		Comment: &comment.Comment{
			Id:        commentModel.ID,
			ArticleId: commentModel.ArticleID,
			Content:   commentModel.Content,
			ParentId:  commentModel.ParentID,
			UserId:    commentModel.UserID,
		},
	}, nil
}
