package logic

import (
	"coderhub/model"
	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"
	"coderhub/shared/SnowFlake"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Upload 上传图片
func (l *UploadLogic) Upload(in *image.UploadRequest) (*image.ImageInfo, error) {
	// TODO: 上传图片到 MinIO

	// 插入数据库
	ID := SnowFlake.GenID()
	err := l.svcCtx.ImageRepository.Create(l.ctx, &model.Image{
		ID:           ID,
		BucketName:   "",
		ObjectName:   "",
		URL:          "",
		ThumbnailURL: "",
		ContentType:  "",
		Size:         0,
		Width:        0,
		Height:       0,
		UploadIP:     "",
		UserID:       in.UserId,
		Status:       model.ImageStatusActive,
	})
	if err != nil {
		return nil, err
	}

	return &image.ImageInfo{
		ImageId:      ID,
		BucketName:   "",
		ObjectName:   "",
		Url:          "",
		ThumbnailUrl: "",
		ContentType:  "",
		Size:         0,
		Width:        0,
		Height:       0,
		UploadIp:     "",
		UserId:       in.UserId,
		Status:       model.ImageStatusActive,
		CreatedAt:    time.Now().Format(time.DateTime),
	}, nil
}
