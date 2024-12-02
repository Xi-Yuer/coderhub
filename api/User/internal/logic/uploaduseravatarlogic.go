package logic

import (
	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"
	"coderhub/conf"
	"coderhub/rpc/User/userservice"
	"coderhub/shared/MetaData"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadUserAvatarLogic) UploadUserAvatar(in *types.UploadUserAvatarRequest) (resp *types.UploadUserAvatarResponse, err error) {
	// 权限验证
	userID, err := MetaData.GetUserID(l.ctx)
	if err != nil {
		return l.errorResp(err)
	}
	ctx := MetaData.SetUserMetaData(l.ctx) // 设置元数据

	// 上传图片
	response, err := l.svcCtx.UserService.UploadAvatar(ctx, &userservice.UploadAvatarRequest{
		UserId:  userID,
		ImageId: in.ImageId,
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
			Url:          data.Url,
			ThumbnailUrl: data.ThumbnailUrl,
			CreatedAt:    data.CreatedAt,
		},
	}, nil
}
