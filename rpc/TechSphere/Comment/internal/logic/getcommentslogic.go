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

type GetCommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentsLogic {
	return &GetCommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetComments 获取评论列表
func (l *GetCommentsLogic) GetComments(in *comment.GetCommentsRequest) (*comment.GetCommentsResponse, error) {
	comments, total, err := l.svcCtx.CommentRepository.ListByArticleID(l.ctx, in.ArticleId, int64(in.Page), int64(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 构建树形结构
	rootComments := l.buildTree(comments)

	return &comment.GetCommentsResponse{
		Comments: rootComments,
		Total:    int32(total),
	}, nil
}

// buildTree 构建树形结构
func (l *GetCommentsLogic) buildTree(comments []model.Comment) []*comment.Comment {
	if len(comments) == 0 {
		return nil
	}

	// 收集所有评论ID
	commentIds := make([]int64, len(comments))
	for i, val := range comments {
		commentIds[i] = val.ID
	}

	l.Logger.Infof("正在获取评论的图片关联，评论IDs: %v", commentIds)

	imageRelations, err := l.svcCtx.ImageRelationService.BatchGetImagesByEntity(l.ctx, &imageRelation.BatchGetImagesByEntityRequest{
		EntityIds:  commentIds,
		EntityType: model.ImageRelationComment,
	})
	if err != nil {
		l.Logger.Errorf("获取评论图片失败: %v", err)
		return make([]*comment.Comment, 0)
	}

	// 构建评论ID到图片列表的映射
	commentImages := make(map[int64][]*comment.CommentImage)
	for _, img := range imageRelations.Relations {
		l.Logger.Infof("处理图片关联: EntityId=%d, ImageId=%d", img.EntityId, img.ImageId)
		// 只有当图片ID大于0时才处理
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
	userIds := make([]int64, len(comments))
	for i, val := range comments {
		userIds[i] = val.UserID
	}
	users, err := l.svcCtx.UserService.BatchGetUserByID(l.ctx, &userservice.BatchGetUserByIDRequest{
		UserIds: userIds,
	})
	if err != nil {
		l.Logger.Errorf("获取用户信息失败: %v", err)
		return make([]*comment.Comment, 0)
	}
	// 构建用户信息映射
	userInfos := make(map[int64]*comment.UserInfo)
	if users != nil && len(users.UserInfos) > 0 {
		for _, user := range users.UserInfos {
			if user != nil {
				l.Logger.Infof("映射用户信息: userId=%d, userName=%s", user.UserId, user.UserName)
				userInfos[user.UserId] = &comment.UserInfo{
					UserId:   user.UserId,
					Username: user.UserName,
					Avatar:   user.Avatar,
				}
			}
		}
	}

	rootComments := make([]*comment.Comment, len(comments))
	for i, val := range comments {
		// 确保每个评论的图片列表都被初始化
		if _, ok := commentImages[val.ID]; !ok {
			commentImages[val.ID] = make([]*comment.CommentImage, 0)
		}

		rootComments[i] = &comment.Comment{
			Id:        val.ID,
			ArticleId: val.ArticleID,
			Content:   val.Content,
			ParentId:  val.ParentID,
			UserInfo:  userInfos[val.UserID],
			Replies:   l.buildTree(val.Replies),
			LikeCount: val.LikeCount,
			Images:    commentImages[val.ID],
			CreatedAt: val.CreatedAt.Unix(),
			UpdatedAt: val.UpdatedAt.Unix(),
		}
	}
	return rootComments
}
