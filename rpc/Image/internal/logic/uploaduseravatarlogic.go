package logic

import (
	"context"

	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传用户头像
func (l *UploadUserAvatarLogic) UploadUserAvatar(in *image.UploadUserAvatarRequest) (*image.UploadUserAvatarResponse, error) {
	// todo: add your logic here and delete this line

	return &image.UploadUserAvatarResponse{}, nil
}
