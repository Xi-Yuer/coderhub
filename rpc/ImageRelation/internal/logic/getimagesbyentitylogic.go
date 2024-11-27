package logic

import (
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"
	"coderhub/rpc/ImageRelation/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImagesByEntityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetImagesByEntityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImagesByEntityLogic {
	return &GetImagesByEntityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetImagesByEntity 获取实体关联的图片列表
func (l *GetImagesByEntityLogic) GetImagesByEntity(in *imageRelation.GetImagesByEntityRequest) (*imageRelation.GetImagesByEntityResponse, error) {
	// 获取实体关联的图片列表
	images, err := l.svcCtx.ImageService.GetImagesByEntity(l.ctx, &imageRelation.GetImagesByEntityRequest{
		EntityId:   in.EntityId,
		EntityType: in.EntityType,
	})
	if err != nil {
		return nil, err
	}
	// 将图片信息转换为ImageInfo
	imageInfos := make([]*imageRelation.ImageInfo, 0, len(images.Images))
	for _, image := range images.Images {
		imageInfos = append(imageInfos, &imageRelation.ImageInfo{
			ImageId:      image.ImageId,
			Url:          image.Url,
			ThumbnailUrl: image.ThumbnailUrl,
			Width:        image.Width,
			Height:       image.Height,
			Sort:         image.Sort,
		})
	}

	return &imageRelation.GetImagesByEntityResponse{
		Images: imageInfos,
	}, nil
}
