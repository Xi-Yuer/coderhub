package logic

import (
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Comment/commentservice"
	"context"

	"coderhub/api/TechSphere/Comment/internal/svc"
	"coderhub/api/TechSphere/Comment/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewCreateCommentLogic 创建评论
func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCommentLogic) CreateComment(req *types.CreateCommentReq) (resp *types.CreateCommentResp, err error) {
	comment, err := l.svcCtx.CommentService.CreateComment(l.ctx, &commentservice.CreateCommentRequest{
		ArticleId: req.ArticleId,
		Content:   req.Content,
		ParentId:  req.ParentId,
		ImageIds:  req.ImageIds,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp(comment)
}

func (l *CreateCommentLogic) successResp(comment *commentservice.CreateCommentResponse) (*types.CreateCommentResp, error) {
	return &types.CreateCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: types.Comment{
			Id: comment.Comment.Id,
		},
	}, nil
}

func (l *CreateCommentLogic) errorResp(err error) (*types.CreateCommentResp, error) {
	return &types.CreateCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
