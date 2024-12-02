package logic

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"coderhub/model"
	"coderhub/rpc/Image/image"
	"coderhub/rpc/ImageRelation/imagerelationservice"
	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"
	"coderhub/shared/MetaData"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UploadAvatar 上传用户头像
func (l *UploadAvatarLogic) UploadAvatar(in *user.UploadAvatarRequest) (*user.UploadAvatarResponse, error) {
	// 权限校验
	var (
		userId string
		err    error
	)
	// 从 metadata 中获取 userId
	if userId, err = MetaData.GetUserMetaData(l.ctx); err != nil {
		return nil, err
	}

	if userId != strconv.FormatInt(in.UserId, 10) {
		return nil, fmt.Errorf("非法操作")
	}

	// 保存图片关系
	createRelationRequest := &imagerelationservice.CreateRelationRequest{
		ImageId:    in.ImageId,
		EntityId:   in.UserId,
		EntityType: model.ImageRelationUserAvatar,
	}
	_, err = l.svcCtx.ImageRelationService.BatchCreateRelation(
		l.ctx,
		&imagerelationservice.BatchCreateRelationRequest{
			Relations: []*imagerelationservice.CreateRelationRequest{
				createRelationRequest,
			},
		},
	)
	if err != nil {
		// 事务回滚
		_, err := l.svcCtx.ImageRelationService.BatchDeleteRelation(l.ctx, &imagerelationservice.BatchDeleteRelationRequest{
			Ids: []int64{in.ImageId},
		})
		if err != nil {
			return nil, err
		}
		l.Errorf("保存图片关系失败: %v", err)
		return nil, err
	}
	// 获取图片信息
	imageInfo, err := l.svcCtx.ImageService.Get(l.ctx, &image.GetRequest{
		ImageId: in.ImageId,
	})
	if err != nil {
		return nil, err
	}
	// 更新用户头像
	userInfo, err := l.svcCtx.UserRepository.GetUserByID(in.UserId)
	if err != nil {
		return nil, err
	}
	l.Logger.Infof("imageRelation用户头像: %v", imageInfo.Url)
	userInfo.Avatar = sql.NullString{String: imageInfo.Url, Valid: true}
	l.Logger.Infof("更新用户头像: %v", userInfo.Avatar)
	err = l.svcCtx.UserRepository.UpdateUser(userInfo)
	if err != nil {
		l.Errorf("更新用户头像失败: %v", err)
		return nil, err
	}

	return &user.UploadAvatarResponse{
		ImageId:      in.ImageId,
		Url:          imageInfo.Url,
		ThumbnailUrl: imageInfo.ThumbnailUrl,
		CreatedAt:    imageInfo.CreatedAt,
	}, nil
}
