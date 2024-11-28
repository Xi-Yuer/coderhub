package logic

import (
	"coderhub/rpc/Image/image"
	"context"

	"coderhub/api/Image/internal/svc"
	"coderhub/api/Image/internal/types"
	"coderhub/conf"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUploadLogic 上传图片
func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadRequest) (resp *types.UploadResponse, err error) {
	response, err := l.svcCtx.ImageService.Upload(l.ctx, &image.UploadRequest{
		File:        req.File,
		Filename:    req.Filename,
		UserId:      req.UserId,
		ContentType: req.ContentType,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp(response)
}
func (l *UploadLogic) successResp(response *image.ImageInfo) (*types.UploadResponse, error) {
	return &types.UploadResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpStatusOK,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: types.ImageInfo{
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
		},
	}, nil
}
func (l *UploadLogic) errorResp(err error) (*types.UploadResponse, error) {
	return &types.UploadResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
