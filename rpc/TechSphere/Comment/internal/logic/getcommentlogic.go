package logic

import (
	"context"

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
			Images:    nil,
		},
	}, nil
}
