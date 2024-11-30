package logic

import (
	"context"
	"strconv"

	"coderhub/model"
	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/TechSphere/Comment/comment"
	"coderhub/rpc/TechSphere/Comment/internal/svc"
	"coderhub/rpc/User/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentRepliesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentRepliesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentRepliesLogic {
	return &GetCommentRepliesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCommentReplies 获取某条评论的子评论列表
func (l *GetCommentRepliesLogic) GetCommentReplies(in *comment.GetCommentRepliesRequest) (*comment.GetCommentRepliesResponse, error) {
	comments, total, err := l.svcCtx.CommentRepository.ListReplies(l.ctx, in.CommentId, int64(in.Page), int64(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 收集所有评论ID
	commentIds := make([]int64, len(comments))
	for i, val := range comments {
		commentIds[i] = val.ID
	}

	l.Logger.Infof("正在获取子评论的图片关联，评论IDs: %v", commentIds)

	imageRelations, err := l.svcCtx.ImageRelationService.BatchGetImagesByEntity(l.ctx, &imageRelation.BatchGetImagesByEntityRequest{
		EntityIds:  commentIds,
		EntityType: model.ImageRelationComment,
	})
	if err != nil {
		l.Logger.Errorf("获取子评论图片失败: %v", err)
		return nil, err
	}

	// 构建评论ID到图片列表的映射
	commentImages := make(map[int64][]*comment.CommentImage)
	for _, img := range imageRelations.Relations {
		l.Logger.Infof("处理图片关联: EntityId=%d, ImageId=%d", img.EntityId, img.ImageId)
		if img.ImageId > 0 {
			imageId := strconv.FormatInt(img.ImageId, 10)
			commentImages[img.EntityId] = append(commentImages[img.EntityId], &comment.CommentImage{
				ImageId:      imageId,
				Url:          img.Url,
				ThumbnailUrl: img.ThumbnailUrl,
			})
		}
	}
	// 获取用户信息
	user, err := l.svcCtx.UserService.GetUserInfo(l.ctx, &userservice.GetUserInfoRequest{
		UserId: comments[0].UserID,
	})
	if err != nil {
		l.Logger.Errorf("获取子评论用户信息失败: %v", err)
		return nil, err
	}

	replies := make([]*comment.Comment, len(comments))
	for i, val := range comments {
		if _, ok := commentImages[val.ID]; !ok {
			commentImages[val.ID] = make([]*comment.CommentImage, 0)
		}

		replies[i] = &comment.Comment{
			Id:        val.ID,
			ArticleId: val.ArticleID,
			Content:   val.Content,	
			ParentId:  val.ParentID,
			UserInfo: &comment.UserInfo{
				UserId:   user.UserId,
				Username: user.UserName,
				Avatar:   user.Avatar,
			},
			CreatedAt: val.CreatedAt.Unix(),
			UpdatedAt: val.UpdatedAt.Unix(),
			Replies:   nil,
			LikeCount: val.LikeCount,
			Images:    commentImages[val.ID],
		}
	}

	return &comment.GetCommentRepliesResponse{Replies: replies, Total: int32(total)}, nil
}
