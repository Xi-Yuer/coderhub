package logic

import (
	"coderhub/shared/MetaData"
	"coderhub/shared/SnowFlake"
	"context"
	"gorm.io/gorm"
	"time"

	"coderhub/model"
	"coderhub/rpc/TechSphere/Comment/comment"
	"coderhub/rpc/TechSphere/Comment/internal/svc"

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
	userID, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return nil, err
	}
	commentModel := &model.Comment{
		ID:        CommentID,
		ArticleID: in.ArticleId,
		Content:   in.Content,
		ParentID:  in.ParentId,
		UserID:    userID,
		LikeCount: 0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		Version:   0,
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
			CreatedAt: commentModel.CreatedAt.Unix(),
			UpdatedAt: commentModel.UpdatedAt.Unix(),
			Replies:   nil,
			LikeCount: 0,
			Images:    nil,
		},
	}, nil
}
