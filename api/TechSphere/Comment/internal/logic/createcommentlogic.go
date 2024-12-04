package logic

import (
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Comment/commentservice"
	"coderhub/shared/MetaData"
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
	ctx := MetaData.SetUserMetaData(l.ctx)
	userID, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return nil, err
	}
	comment, err := l.svcCtx.CommentService.CreateComment(ctx, &commentservice.CreateCommentRequest{
		ArticleId:  req.ArticleId,
		Content:    req.Content,
		ParentId:   req.ParentId,
		RootId:     req.RootId,
		UserId:     userID,
		ReplyToUid: req.ReplyToUID,
		ImageIds:   req.ImageIds,
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
		Data: &types.Comment{
			Id:           comment.Comment.Id,
			ArticleId:    comment.Comment.ArticleId,
			Content:      comment.Comment.Content,
			ParentId:     comment.Comment.ParentId,
			RootId:       comment.Comment.RootId,
			UserInfo:     &types.UserInfo{UserId: comment.Comment.UserInfo.UserId, Username: comment.Comment.UserInfo.Username, Avatar: comment.Comment.UserInfo.Avatar},
			CreatedAt:    comment.Comment.CreatedAt,
			UpdatedAt:    comment.Comment.UpdatedAt,
			Replies:      nil,
			RepliesCount: 0,
			LikeCount:    0,
			Images:       nil,
		},
	}, nil
}

func (l *CreateCommentLogic) errorResp(err error) (*types.CreateCommentResp, error) {
	return &types.CreateCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
		Data: nil,
	}, nil
}
