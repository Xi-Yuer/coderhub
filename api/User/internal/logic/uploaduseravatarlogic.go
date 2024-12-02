package logic

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"

	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"
	"coderhub/conf"
	"coderhub/rpc/User/userservice"
	"coderhub/shared/MetaData"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadUserAvatarLogic(ctx context.Context, r *http.Request, svcCtx *svc.ServiceContext) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadUserAvatarLogic) UploadUserAvatar() (resp *types.UploadUserAvatarResponse, err error) {
	// 权限验证
	userID, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err)
	}
	ctx := MetaData.SetUserMetaData(l.ctx) // 设置元数据

	err = l.r.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, err
	}
	file, handler, err := l.r.FormFile("file")
	if err != nil {
		return l.errorResp(err)
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			l.Logger.Error("关闭文件失败", err)
		}
	}(file)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return l.errorResp(err)
	}

	// 上传图片
	response, err := l.svcCtx.UserService.UploadAvatar(ctx, &userservice.UploadAvatarRequest{
		File:        fileBytes,
		Filename:    handler.Filename,
		UserId:      userID,
		ContentType: handler.Header.Get("Content-Type"),
	})
	if err != nil {
		return l.errorResp(err)
	}
	return l.successResp(response)
}

func (l *UploadUserAvatarLogic) errorResp(err error) (*types.UploadUserAvatarResponse, error) {
	return &types.UploadUserAvatarResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: err.Error(),
		},
		Data: nil,
	}, nil
}

func (l *UploadUserAvatarLogic) successResp(data *userservice.UploadAvatarResponse) (*types.UploadUserAvatarResponse, error) {
	return &types.UploadUserAvatarResponse{
		Response: types.Response{
			Code:    conf.HttpCode.HttpBadRequest,
			Message: conf.HttpMessage.MsgOK,
		},
		Data: &types.ImageInfo{
			ImageId:      data.ImageId,
			BucketName:   data.BucketName,
			ObjectName:   data.ObjectName,
			Url:          data.Url,
			ThumbnailUrl: data.ThumbnailUrl,
			ContentType:  data.ContentType,
			Size:         data.Size,
			Width:        data.Width,
			Height:       data.Height,
			UploadIp:     data.UploadIp,
			UserId:       data.UserId,
			Status:       data.Status,
			CreatedAt:    data.CreatedAt,
		},
	}, nil
}
