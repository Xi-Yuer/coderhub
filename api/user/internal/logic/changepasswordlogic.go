package logic

import (
	"coderhub/api/user/internal/types"
	"coderhub/rpc/user/user"
	"context"

	"coderhub/api/user/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordRequest) (resp *types.ChangePasswordResponse, err error) {
	// 获取用户信息
	userInfo, err := l.svcCtx.UserService.GetUserInfo(l.ctx, &user.GetUserInfoRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return &types.ChangePasswordResponse{
			Response: types.Response{
				Code:    -1,
				Message: err.Error(),
			},
			Data: false,
		}, nil
	}

	// 检查用户ID是否匹配
	if userInfo.UserId != req.UserId {
		return &types.ChangePasswordResponse{
			Response: types.Response{
				Code:    -1,
				Message: "非法操作",
			},
			Data: false,
		}, nil
	}

	// 更新密码
	_, err = l.svcCtx.UserService.ChangePassword(l.ctx, &user.ChangePasswordRequest{
		UserId:      req.UserId,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return &types.ChangePasswordResponse{
			Response: types.Response{
				Code:    -1,
				Message: err.Error(),
			},
			Data: false,
		}, nil
	}

	// 返回成功响应
	return &types.ChangePasswordResponse{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		Data: true,
	}, nil
}
