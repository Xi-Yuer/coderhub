package commentservicelogic

import (
	"context"

	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/svc"

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
func (l *GetCommentRepliesLogic) GetCommentReplies(in *coderhub.GetCommentRepliesRequest) (*coderhub.GetCommentRepliesResponse, error) {
	// todo: add your logic here and delete this line

	return &coderhub.GetCommentRepliesResponse{}, nil
}
