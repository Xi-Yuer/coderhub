package logic

import (
	"context"

	"coderhub/api/TechSphere/Comment/internal/svc"
	"coderhub/api/TechSphere/Comment/internal/types"
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Comment/commentservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetCommentsLogic 获取评论列表
func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentsLogic) GetComments(req *types.GetCommentsReq) (resp *types.GetCommentsResp, err error) {
	comments, err := l.svcCtx.CommentService.GetComments(l.ctx, &commentservice.GetCommentsRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp(comments)
}

func (l *GetCommentsLogic) successResp(comments *commentservice.GetCommentsResponse) (*types.GetCommentsResp, error) {
	// 创建一个map用于存储所有评论
	commentMap := make(map[int64]*types.Comment)
	// 用于存储顶级评论
	var rootComments []*types.Comment

	// 第一步：将所有评论放入map中
	for _, comment := range comments.Comments {
		commentObj := &types.Comment{
			Id:        comment.Id,
			ArticleId: comment.ArticleId,
			Content:   comment.Content,
			ParentId:  comment.ParentId,
			UserId:    comment.UserId,
			Replies:   make([]*types.Comment, 0),
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
		commentMap[comment.Id] = commentObj

		// 如果是顶级评论（ParentId为0），加入rootComments
		if comment.ParentId == 0 {
			rootComments = append(rootComments, commentObj)
		}
	}

	// 第二步：建立父子关系
	for _, comment := range comments.Comments {
		if comment.ParentId != 0 {
			// 找到父评论，将当前评论添加到父评论的replies中
			if parent, exists := commentMap[comment.ParentId]; exists {
				parent.Replies = append(parent.Replies, commentMap[comment.Id])
			}
		}
	}

	return &types.GetCommentsResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: types.List{
			Comments: rootComments, // 只返回顶级评论，子评论在replies字段中
			Total:    comments.Total,
		},
	}, nil
}

func (l *GetCommentsLogic) errorResp(err error) (*types.GetCommentsResp, error) {
	return &types.GetCommentsResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
