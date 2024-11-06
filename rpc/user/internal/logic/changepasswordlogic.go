package logic

import (
	"coderhub/model"
	"coderhub/shared/bcryptUtil"
	"context"
	"errors"

	"coderhub/rpc/user/internal/svc"
	"coderhub/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ChangePassword 修改密码
func (l *ChangePasswordLogic) ChangePassword(in *user.ChangePasswordRequest) (*user.ChangePasswordResponse, error) {
	var userInfo model.User
	l.svcCtx.SqlDB.First(&userInfo, "id = ?", in.UserId)

	// 验证旧密码是否正确
	if !bcryptUtil.CompareHashAndPassword(userInfo.Password, in.OldPassword) {
		return nil, errors.New("旧密码不正确")
	}

	// 对新密码进行哈希处理
	hashedNewPassword, err := bcryptUtil.PasswordHash(in.NewPassword)
	if err != nil {
		return nil, err
	}

	// 更新用户密码
	if tx := l.svcCtx.SqlDB.Model(&userInfo).Update("password", hashedNewPassword); tx.Error != nil {
		return nil, tx.Error
	}

	// 返回成功响应
	return &user.ChangePasswordResponse{
		Success: true,
	}, nil
}
