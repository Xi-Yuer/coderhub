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

// GetCommentReplies 获取评论回复列表
func (l *GetCommentRepliesLogic) GetCommentReplies(in *comment.GetCommentRepliesRequest) (*comment.GetCommentRepliesResponse, error) {
	// 获取回复列表
	replies, total, err := l.svcCtx.CommentRepository.ListReplies(l.ctx, in.CommentId, int64(in.Page), int64(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 如果没有回复，直接返回空结果
	if len(replies) == 0 {
		return &comment.GetCommentRepliesResponse{
			Replies: make([]*comment.Comment, 0),
			Total:   0,
		}, nil
	}

	// 收集所有回复ID
	replyIds := make([]int64, len(replies))
	userIds := make([]int64, len(replies))
	for i, reply := range replies {
		replyIds[i] = reply.ID
		userIds[i] = reply.UserID
	}

	// 获取图片关联
	imageRelations, err := l.svcCtx.ImageRelationService.BatchGetImagesByEntity(l.ctx, &imageRelation.BatchGetImagesByEntityRequest{
		EntityIds:  replyIds,
		EntityType: model.ImageRelationComment,
	})
	if err != nil {
		l.Logger.Errorf("获取回复图片失败: %v", err)
		return nil, err
	}

	// 构建回复ID到图片列表的映射
	replyImages := make(map[int64][]*comment.CommentImage)
	for _, img := range imageRelations.Relations {
		if img.ImageId > 0 {
			imageId := strconv.FormatInt(img.ImageId, 10)
			replyImages[img.EntityId] = append(replyImages[img.EntityId], &comment.CommentImage{
				ImageId:      imageId,
				Url:          img.Url,
				ThumbnailUrl: img.ThumbnailUrl,
			})
		}
	}

	// 获取用户信息
	users, err := l.svcCtx.UserService.BatchGetUserByID(l.ctx, &userservice.BatchGetUserByIDRequest{
		UserIds: userIds,
	})
	if err != nil {
		l.Logger.Errorf("获取用户信息失败: %v", err)
		return nil, err
	}

	// 添加调试日志
	l.Logger.Infof("用户ID列表: %v", userIds)
	l.Logger.Infof("获取到的用户信息: %+v", users)

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

	// 构建回复列表
	commentReplies := make([]*comment.Comment, len(replies))
	for i, reply := range replies {
		// 确保每个回复的图片列表被初始化
		if _, ok := replyImages[reply.ID]; !ok {
			replyImages[reply.ID] = make([]*comment.CommentImage, 0)
		}
		commentReplies[i] = &comment.Comment{
			Id:        reply.ID,
			ArticleId: reply.ArticleID,
			Content:   reply.Content,
			ParentId:  reply.ParentID,
			UserInfo:  userInfos[reply.UserID],
			LikeCount: reply.LikeCount,
			Images:    replyImages[reply.ID],
			CreatedAt: reply.CreatedAt.Unix(),
			UpdatedAt: reply.UpdatedAt.Unix(),
		}
	}

	return &comment.GetCommentRepliesResponse{
		Replies: commentReplies,
		Total:   int32(total),
	}, nil
}
