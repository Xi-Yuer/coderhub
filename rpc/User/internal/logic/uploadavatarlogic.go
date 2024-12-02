package logic

import (
	"context"

	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"

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

// 上传用户头像
func (l *UploadAvatarLogic) UploadAvatar(in *user.UploadAvatarRequest) (*user.UploadAvatarResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UploadAvatarResponse{}, nil
}
