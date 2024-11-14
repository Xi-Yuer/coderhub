package logic

import (
	"coderhub/model"
	"coderhub/rpc/User/internal/svc"
	"coderhub/rpc/User/user"
	"coderhub/shared/BcryptUtil"
	"coderhub/shared/MetaData"
	"coderhub/shared/Validator"
	"context"
	"errors"
	"fmt"
	"strconv"

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
	if err := Validator.New().Password(in.OldPassword).Password(in.NewPassword).Check(); err != nil {
		return nil, err
	}

	var (
		userId string
		err    error
	)
	// 从 metadata 中获取 userId
	if userId, err = MetaData.GetUserMetaData(l.ctx); err != nil {
		return nil, err
	}

	if userId != strconv.FormatInt(in.UserId, 10) {
		return nil, fmt.Errorf("非法操作")
	}

	userInfo, err := NewGetUserInfoLogic(l.ctx, l.svcCtx).GetUserInfo(&user.GetUserInfoRequest{UserId: in.UserId})
	if err != nil {
		return nil, err
	}

	// 验证旧密码是否正确
	if !BcryptUtil.CompareHashAndPassword(userInfo.Password, in.OldPassword) {
		return nil, errors.New("旧密码不正确")
	}

	// 对新密码进行哈希处理
	hashedNewPassword, err := BcryptUtil.PasswordHash(in.NewPassword)
	if err != nil {
		return nil, err
	}

	// 更新用户密码
	if err := l.svcCtx.UserRepository.UpdateUser(&model.User{ID: in.UserId, Password: hashedNewPassword}); err != nil {
		return nil, err
	}

	// 返回成功响应
	return &user.ChangePasswordResponse{
		Success: true,
	}, nil
}
