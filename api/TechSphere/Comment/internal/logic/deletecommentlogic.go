package logic

import (
	"coderhub/rpc/TechSphere/Comment/commentservice"
	"coderhub/shared/MetaData"
	"context"

	"coderhub/api/TechSphere/Comment/internal/svc"
	"coderhub/api/TechSphere/Comment/internal/types"
	"coderhub/conf"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewDeleteCommentLogic 删除评论
func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCommentLogic) DeleteComment(req *types.DeleteCommentReq) (resp *types.DeleteCommentResp, err error) {
	// 权限校验
	_, err = MetaData.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err)
	}
	ctx := MetaData.SetUserMetaData(l.ctx) // 设置元数据

	_, err = l.svcCtx.CommentService.DeleteComment(ctx, &commentservice.DeleteCommentRequest{
		CommentId: req.CommentId,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp()
}

func (l *DeleteCommentLogic) successResp() (*types.DeleteCommentResp, error) {
	return &types.DeleteCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
	}, nil
}

func (l *DeleteCommentLogic) errorResp(err error) (*types.DeleteCommentResp, error) {
	return &types.DeleteCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
