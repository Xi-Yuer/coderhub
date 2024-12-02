package logic

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"coderhub/model"
	"coderhub/rpc/Image/imageservice"
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

	// 上传图片
	response, err := l.svcCtx.ImageService.Upload(l.ctx, &imageservice.UploadRequest{
		File:        in.File,
		Filename:    in.Filename,
		UserId:      in.UserId,
		ContentType: in.ContentType,
	})
	if err != nil {
		return nil, err
	}

	// 保存图片关系
	imageRelation, err := l.svcCtx.ImageRelationService.CreateRelation(l.ctx, &imagerelationservice.CreateRelationRequest{
		ImageId:    response.ImageId,
		EntityId:   in.UserId,
		EntityType: model.ImageRelationUserAvatar,
		Sort:       0,
	})
	if err != nil {
		// 事务回滚
		_, err := l.svcCtx.ImageRelationService.BatchDeleteRelation(l.ctx, &imagerelationservice.BatchDeleteRelationRequest{
			Ids: []int64{response.ImageId},
		})
		if err != nil {
			return nil, err
		}
		l.Errorf("保存图片关系失败: %v", err)
		return nil, err
	}
	// 更新用户头像
	// 获取用户
	userInfo, err := l.svcCtx.UserRepository.GetUserByID(in.UserId)
	if err != nil {
		return nil, err
	}
	userInfo.Avatar = sql.NullString{String: imageRelation.Relation.Url, Valid: true}
	err = l.svcCtx.UserRepository.UpdateUser(userInfo)
	if err != nil {
		l.Errorf("更新用户头像失败: %v", err)
		return nil, err
	}

	return &user.UploadAvatarResponse{
		ImageId:      response.ImageId,
		BucketName:   response.BucketName,
		ObjectName:   response.ObjectName,
		Url:          response.Url,
		ThumbnailUrl: response.ThumbnailUrl,
		ContentType:  response.ContentType,
		Size:         response.Size,
		Width:        response.Width,
		Height:       response.Height,
		UploadIp:     response.UploadIp,
		UserId:       response.UserId,
		Status:       response.Status,
		CreatedAt:    response.CreatedAt,
	}, nil
}
