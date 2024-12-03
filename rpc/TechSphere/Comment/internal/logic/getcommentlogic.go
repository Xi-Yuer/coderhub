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

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetComment 获取单个评论详情
func (l *GetCommentLogic) GetComment(in *comment.GetCommentRequest) (*comment.GetCommentResponse, error) {
	commentModel, err := l.svcCtx.CommentRepository.GetByID(l.ctx, in.CommentId)
	if err != nil {
		return nil, err
	}
	// 评论ID
	commentId := []int64{commentModel.ID}
	// 获取与评论关联的所有图片
	imageRelations, err := l.svcCtx.ImageRelationService.BatchGetImagesByEntity(l.ctx, &imageRelation.BatchGetImagesByEntityRequest{
		EntityIds:  commentId,
		EntityType: model.ImageRelationComment,
	})
	if err != nil {
		return nil, err
	}
	// 获取用户信息
	user, err := l.svcCtx.UserService.GetUserInfo(l.ctx, &userservice.GetUserInfoRequest{
		UserId: commentModel.UserID,
	})
	if err != nil {
		return nil, err
	}

	images := make([]*comment.CommentImage, 0)
	// 将图片关联转换为评论图片
	for _, val := range imageRelations.Relations {
		imageId := strconv.FormatInt(val.ImageId, 10)
		images = append(images, &comment.CommentImage{
			ImageId:      imageId,
			Url:          val.Url,
			ThumbnailUrl: val.ThumbnailUrl,
		})
	}
	// 获取评论点赞数
	likeCount, err := l.svcCtx.CommentRelationLikeRepository.List(l.ctx, commentModel.ID)
	if err != nil {
		return nil, err
	}
	return &comment.GetCommentResponse{
		Comment: &comment.Comment{
			Id:        commentModel.ID,
			ArticleId: commentModel.ArticleID,
			Content:   commentModel.Content,
			ParentId:  commentModel.ParentID,
			UserInfo: &comment.UserInfo{
				UserId:   user.UserId,
				Username: user.UserName,
				Avatar:   user.Avatar,
			},
			CreatedAt: commentModel.CreatedAt.Unix(),
			UpdatedAt: commentModel.UpdatedAt.Unix(),
			LikeCount: int32(likeCount),
			Images:    images,
		},
	}, nil
}
