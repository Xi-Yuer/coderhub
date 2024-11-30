package logic

import (
	"context"

	"coderhub/api/TechSphere/Comment/internal/svc"
	"coderhub/api/TechSphere/Comment/internal/types"
	"coderhub/conf"
	"coderhub/rpc/TechSphere/Comment/commentservice"
	"coderhub/rpc/User/userservice"

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
	// 获取图片
	images := make([]types.CommentImage, len(comment.Comment.Images))
	for i, image := range comment.Comment.Images {
		images[i] = types.CommentImage{
			ImageId:      image.ImageId,
			Url:          image.Url,
			ThumbnailUrl: image.ThumbnailUrl,
		}
	}
	// 获取用户信息
	user, err := l.svcCtx.UserService.GetUserInfo(l.ctx, &userservice.GetUserInfoRequest{
		UserId: comment.Comment.UserInfo.UserId,
	})
	if err != nil {
		return l.errorResp(err)
	}
	userInfo := types.UserInfo{
		UserId:   user.UserId,
		Username: user.UserName,
		Avatar:   user.Avatar,
	}
	return &types.GetCommentResp{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: types.Comment{
			Id:        comment.Comment.Id,
			ArticleId: comment.Comment.ArticleId,
			Content:   comment.Comment.Content,
			ParentId:  comment.Comment.ParentId,
			UserInfo:  userInfo,
			CreatedAt: comment.Comment.CreatedAt,
			UpdatedAt: comment.Comment.UpdatedAt,
			Replies:   nil,
			LikeCount: comment.Comment.LikeCount,
			Images:    images,
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
