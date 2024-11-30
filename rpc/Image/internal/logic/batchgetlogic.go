package logic

import (
	"context"
	"time"

	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchGetLogic {
	return &BatchGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BatchGet 批量获取图片信息
func (l *BatchGetLogic) BatchGet(in *image.BatchGetRequest) (*image.BatchGetResponse, error) {
	images, err := l.svcCtx.ImageRepository.BatchGetImagesByID(l.ctx, in.ImageIds)
	if err != nil {
		return nil, err
	}
	imageInfos := make([]*image.ImageInfo, 0)
	for _, value := range images {
		imageInfos = append(imageInfos, &image.ImageInfo{
			ImageId:      value.ID,
			BucketName:   value.BucketName,
			ObjectName:   value.ObjectName,
			Url:          value.URL,
			ThumbnailUrl: value.ThumbnailURL,
			ContentType:  value.ContentType,
			Size:         value.Size,
			Width:        value.Width,
			Height:       value.Height,
			UploadIp:     value.UploadIP,
			UserId:       value.UserID,
			Status:       value.Status,
			CreatedAt:    value.CreatedAt.Format(time.DateTime),
		})
	}

	return &image.BatchGetResponse{
		Images: imageInfos,
	}, nil
}
