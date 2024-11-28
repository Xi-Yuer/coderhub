package logic

import (
	"context"

	"coderhub/model"
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
			Images:    nil,
			CreatedAt: val.CreatedAt.Unix(),
			UpdatedAt: val.UpdatedAt.Unix(),
		}
	}
	return rootComments
}
