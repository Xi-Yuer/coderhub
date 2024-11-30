package logic

import (
	"context"

	"coderhub/rpc/TechSphere/Comment/comment"
	"coderhub/rpc/TechSphere/Comment/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentRepliesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentRepliesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentRepliesLogic {
	return &GetCommentRepliesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取某条评论的子评论列表
func (l *GetCommentRepliesLogic) GetCommentReplies(in *comment.GetCommentRepliesRequest) (*comment.GetCommentRepliesResponse, error) {
	// todo: add your logic here and delete this line

	return &comment.GetCommentRepliesResponse{}, nil
}
