package logic

import (
	"context"

	"coderhub/rpc/TechSphere/Comment/comment"
	"coderhub/rpc/TechSphere/Comment/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLikeCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLikeCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLikeCountLogic {
	return &UpdateCommentLikeCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新评论点赞数
func (l *UpdateCommentLikeCountLogic) UpdateCommentLikeCount(in *comment.UpdateCommentLikeCountRequest) (*comment.UpdateCommentLikeCountResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.UpdateCommentLikeCountResponse{}, nil
}
