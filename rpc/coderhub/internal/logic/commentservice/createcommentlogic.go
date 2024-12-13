package commentservicelogic

import (
	"coderhub/model"
	imagerelationservicelogic "coderhub/rpc/coderhub/internal/logic/imagerelationservice"
	userservicelogic "coderhub/rpc/coderhub/internal/logic/userservice"
	"coderhub/shared/utils"
	"context"
	"strconv"

	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateComment 创建评论
func (l *CreateCommentLogic) CreateComment(in *coderhub.CreateCommentRequest) (*coderhub.CreateCommentResponse, error) {
	CommentID := utils.GenID()
	commentModel := &model.Comment{
		ID:         CommentID,
		ArticleID:  in.ArticleId,
		Content:    in.Content,
		ParentID:   in.ParentId,
		RootID:     in.RootId,
		UserID:     in.UserId,
		ReplyToUID: in.ReplyToUid,
	}
	// 获取用户信息
	userService := userservicelogic.NewGetUserInfoLogic(l.ctx, l.svcCtx)
	user, err := userService.GetUserInfo(&coderhub.GetUserInfoRequest{
		UserId: in.UserId,
	})
	// 如果评论有携带图片，则需要创建图片关联
	var imageRelationModels []*coderhub.CreateRelationRequest
	if len(in.ImageIds) > 0 {
		for _, imageId := range in.ImageIds {
			imageIdInt, err := strconv.ParseInt(imageId, 10, 64)
			if err != nil {
				return nil, err
			}
			imageRelationModels = append(imageRelationModels, &coderhub.CreateRelationRequest{
				ImageId:    imageIdInt,
				EntityId:   CommentID,
				EntityType: model.ImageRelationComment,
			})
		}
	}
	imageBatchCreateService := imagerelationservicelogic.NewBatchCreateRelationLogic(l.ctx, l.svcCtx)
	_, err = imageBatchCreateService.BatchCreateRelation(&coderhub.BatchCreateRelationRequest{
		Relations: imageRelationModels,
	})
	if err != nil {
		return nil, err
	}
	// 创建评论
	if err := l.svcCtx.CommentRepository.Create(l.ctx, commentModel); err != nil {
		// 事务回滚
		imageBatchDeleteService := imagerelationservicelogic.NewBatchDeleteRelationLogic(l.ctx, l.svcCtx)
		_, _ = imageBatchDeleteService.BatchDeleteRelation(&coderhub.BatchDeleteRelationRequest{
			Ids: []int64{CommentID},
		})
		return nil, err
	}
	return &coderhub.CreateCommentResponse{
		Comment: &coderhub.Comment{
			Id:           commentModel.ID,
			ArticleId:    commentModel.ArticleID,
			Content:      commentModel.Content,
			ParentId:     commentModel.ParentID,
			RootId:       commentModel.RootID,
			UserInfo:     &coderhub.CommentUserInfo{UserId: user.UserId, Username: user.UserName, Avatar: user.Avatar},
			CreatedAt:    commentModel.CreatedAt.Unix(),
			UpdatedAt:    commentModel.UpdatedAt.Unix(),
			Replies:      nil,
			RepliesCount: 0,
			LikeCount:    0,
			Images:       nil,
		},
	}, nil
}
