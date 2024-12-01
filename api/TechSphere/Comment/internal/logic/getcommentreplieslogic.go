package logic

import (
	"context"

	"coderhub/api/TechSphere/Comment/internal/svc"
	"coderhub/api/TechSphere/Comment/internal/types"
	"coderhub/rpc/TechSphere/Comment/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentRepliesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取某条评论的子评论列表
func NewGetCommentRepliesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentRepliesLogic {
	return &GetCommentRepliesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentRepliesLogic) GetCommentReplies(req *types.GetCommentRepliesReq) (resp *types.GetCommentRepliesResp, err error) {
	reply, err := l.svcCtx.CommentService.GetCommentReplies(l.ctx, &comment.GetCommentRepliesRequest{
		CommentId: req.CommentId,
		Page:      int32(req.Page),
		PageSize:  int32(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	replies := make([]*types.Comment, len(reply.Replies))
	for i, val := range reply.Replies {
		images := make([]types.CommentImage, len(val.Images))
		for j, img := range val.Images {
			images[j] = types.CommentImage{
				ImageId:      img.ImageId,
				Url:          img.Url,
				ThumbnailUrl: img.ThumbnailUrl,
			}
		}
		replies[i] = &types.Comment{
			Id:        val.Id,
			ArticleId: val.ArticleId,
			Content:   val.Content,
			ParentId:  val.ParentId,
			UserInfo: types.UserInfo{
				UserId:   val.UserInfo.UserId,
				Username: val.UserInfo.Username,
				Avatar:   val.UserInfo.Avatar,
			},
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			Replies:   nil,
			LikeCount: val.LikeCount,
			Images:    images,
		}
	}

	return &types.GetCommentRepliesResp{
		Response: types.Response{},
		Data: types.List{
			List:  replies,
			Total: reply.Total,
		},
	}, nil
}