package logic

import (
	"context"

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
	commentModels, total, err := l.svcCtx.CommentRepository.ListByArticleID(l.ctx, in.ArticleId, int64(in.Page), int64(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 只保留一次评论转换逻辑
	comments := make([]*comment.Comment, len(commentModels))
	commentMap := make(map[int64]*comment.Comment)
	for i, model := range commentModels {
		comments[i] = &comment.Comment{
			Id:        model.ID,
			ArticleId: model.ArticleID,
			Content:   model.Content,
			ParentId:  model.ParentID,
			UserId:    model.UserID,
			Replies:   make([]*comment.Comment, 0),
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix(),
		}
		commentMap[comments[i].Id] = comments[i]
	}

	// 构建评论树
	var rootComments []*comment.Comment

	// 递归构建评论树的辅助函数
	var buildCommentTree func(parentId int64) []*comment.Comment
	buildCommentTree = func(parentId int64) []*comment.Comment {
		var children []*comment.Comment
		for _, c := range comments {
			if c.ParentId == parentId {
				// 递归获取子评论
				c.Replies = buildCommentTree(c.Id)
				children = append(children, c)
			}
		}
		return children
	}

	// 获取所有顶级评论（ParentId = 0）及其所有子评论
	rootComments = buildCommentTree(0)

	return &comment.GetCommentsResponse{
		Comments: rootComments,
		Total:    int32(total),
	}, nil
}
