package logic

import (
	"coderhub/shared/utils"
	"context"
	"fmt"
	"strconv"

	"coderhub/rpc/Image/image"
	"coderhub/rpc/Image/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除图片
func (l *DeleteLogic) Delete(in *image.DeleteRequest) (*image.DeleteResponse, error) {
	// 权限校验
	var (
		userId string
		err    error
	)
	if userId, err = utils.GetUserMetaData(l.ctx); err != nil {
		return nil, err
	}

	if userId != strconv.FormatInt(in.UserId, 10) {
		return nil, fmt.Errorf("非法操作")
	}

	err = l.svcCtx.ImageRepository.Delete(l.ctx, in.ImageId)
	if err != nil {
		return nil, err
	}
	return &image.DeleteResponse{
		Success: true,
	}, nil
}
