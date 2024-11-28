package logic

import (
	"context"

	"coderhub/api/Image/internal/svc"
	"coderhub/api/Image/internal/types"
	"coderhub/conf"
	"coderhub/rpc/Image/image"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewGetLogic 获取图片信息
func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.GetRequest) (resp *types.GetResponse, err error) {
	response, err := l.svcCtx.ImageService.Get(l.ctx, &image.GetRequest{
		ImageId: req.ImageId,
	})
	if err != nil {
		return l.errorResp(err)
	}

	return l.successResp(response)
}

func (l *GetLogic) successResp(response *image.ImageInfo) (*types.GetResponse, error) {
	return &types.GetResponse{
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

func (l *GetLogic) errorResp(err error) (*types.GetResponse, error) {
	return &types.GetResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
	}, nil
}
