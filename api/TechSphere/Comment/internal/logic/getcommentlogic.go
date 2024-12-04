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

// NewGetCommentLogic 获取单个评论
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
	if comment == nil {
		return &types.GetCommentResp{
			Response: types.Response{
				Code:    conf.HttpCode.HttpNotFound,
				Message: conf.HttpMessage.MsgFailed,
			},
		}, nil
	}

	// 构建用户信息，处理空值情况
	var userInfo *types.UserInfo
	if comment.Comment.UserInfo != nil {
		userInfo = &types.UserInfo{
			UserId:   comment.Comment.UserInfo.UserId,
			Username: comment.Comment.UserInfo.Username,
			Avatar:   comment.Comment.UserInfo.Avatar,
		}
	}

	var replyToUserInfo *types.UserInfo
	if comment.Comment.ReplyToUserInfo != nil {
		replyToUserInfo = &types.UserInfo{
			UserId:   comment.Comment.ReplyToUserInfo.UserId,
			Username: comment.Comment.ReplyToUserInfo.Username,
			Avatar:   comment.Comment.ReplyToUserInfo.Avatar,
		}
	}

	// 获取图片
	images := make([]types.CommentImage, len(comment.Comment.Images))
	for i, image := range comment.Comment.Images {
		images[i] = types.CommentImage{
			ImageId:      image.ImageId,
			Url:          image.Url,
			ThumbnailUrl: image.ThumbnailUrl,
		}
	}

	return &types.GetCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: &types.Comment{
			Id:              comment.Comment.Id,
			ArticleId:       comment.Comment.ArticleId,
			Content:         comment.Comment.Content,
			ParentId:        comment.Comment.ParentId,
			RootId:          comment.Comment.RootId,
			UserInfo:        userInfo,
			CreatedAt:       comment.Comment.CreatedAt,
			UpdatedAt:       comment.Comment.UpdatedAt,
			Replies:         nil,
			RepliesCount:    comment.Comment.RepliesCount,
			ReplyToUserInfo: replyToUserInfo,
			LikeCount:       comment.Comment.LikeCount,
			Images:          images,
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
