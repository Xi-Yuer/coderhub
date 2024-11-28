package logic

import (
	"coderhub/model"
	"coderhub/rpc/TechSphere/Comment/comment"
	"coderhub/rpc/TechSphere/Comment/internal/svc"
	"coderhub/shared/SnowFlake"
	"context"
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
	// 创建评论
	if err := l.svcCtx.CommentRepository.Create(l.ctx, commentModel); err != nil {
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
