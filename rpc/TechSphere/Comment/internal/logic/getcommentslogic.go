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
	comments := make([]*comment.Comment, len(commentModels))
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
	}

	// 构建评论树
	commentMap := make(map[int64]*comment.Comment)
	var rootComments []*comment.Comment

	// 第一步：将所有评论放入map中
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

	// 第二步：构建树形结构
	for _, c := range comments {
		if c.ParentId == 0 {
			// 根评论
			rootComments = append(rootComments, c)
		} else {
			// 子评论
			if parent, exists := commentMap[c.ParentId]; exists {
				parent.Replies = append(parent.Replies, c)
			}
		}
	}

	return &comment.GetCommentsResponse{
		Comments: rootComments,
		Total:    int32(total),
	}, nil
}
