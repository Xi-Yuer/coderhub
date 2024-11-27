package logic

import (
	"context"

	"coderhub/api/TechSphere/Comment/internal/svc"
	"coderhub/api/TechSphere/Comment/internal/types"
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Comment/commentservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取单个评论
func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentLogic) GetComment(req *types.GetCommentReq) (resp *types.GetCommentResp, err error) {
	comment, err := l.svcCtx.CommentService.GetComment(l.ctx, &commentservice.GetCommentRequest{
		CommentId: req.CommentId,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp(comment)
}

func (l *GetCommentLogic) successResp(comment *commentservice.GetCommentResponse) (*types.GetCommentResp, error) {
	return &types.GetCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: types.Comment{
			Id: comment.Comment.Id,
		},
	}, nil
}

func (l *GetCommentLogic) errorResp(err error) (*types.GetCommentResp, error) {
	return &types.GetCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
