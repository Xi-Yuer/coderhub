// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package userservice

import (
	"context"

	"coderhub/rpc/user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthorizeRequest             = user.AuthorizeRequest
	AuthorizeResponse            = user.AuthorizeResponse
	ChangePasswordRequest        = user.ChangePasswordRequest
	ChangePasswordResponse       = user.ChangePasswordResponse
	CheckUserExistsRequest       = user.CheckUserExistsRequest
	CheckUserExistsResponse      = user.CheckUserExistsResponse
	CreateUserRequest            = user.CreateUserRequest
	CreateUserResponse           = user.CreateUserResponse
	DeleteUserRequest            = user.DeleteUserRequest
	DeleteUserResponse           = user.DeleteUserResponse
	GenerateTokenRequest         = user.GenerateTokenRequest
	GenerateTokenResponse        = user.GenerateTokenResponse
	GetUserInfoByUsernameRequest = user.GetUserInfoByUsernameRequest
	GetUserInfoRequest           = user.GetUserInfoRequest
	GetUserInfoResponse          = user.GetUserInfoResponse
	ResetPasswordRequest         = user.ResetPasswordRequest
	ResetPasswordResponse        = user.ResetPasswordResponse
	UpdateUserInfoRequest        = user.UpdateUserInfoRequest
	UpdateUserInfoResponse       = user.UpdateUserInfoResponse

	UserService interface {
		// 授权
		Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
		// 检查用户是否存在
		CheckUserExists(ctx context.Context, in *CheckUserExistsRequest, opts ...grpc.CallOption) (*CheckUserExistsResponse, error)
		// 创建用户
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		// 获取用户信息
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
		GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
		// 更新用户信息
		UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
		// 修改密码
		ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
		// 重置密码
		ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
		// 删除用户
		DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

// 授权
func (m *defaultUserService) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Authorize(ctx, in, opts...)
}

// 检查用户是否存在
func (m *defaultUserService) CheckUserExists(ctx context.Context, in *CheckUserExistsRequest, opts ...grpc.CallOption) (*CheckUserExistsResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.CheckUserExists(ctx, in, opts...)
}

// 创建用户
func (m *defaultUserService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

// 获取用户信息
func (m *defaultUserService) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserService) GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfoByUsername(ctx, in, opts...)
}

// 更新用户信息
func (m *defaultUserService) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

// 修改密码
func (m *defaultUserService) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.ChangePassword(ctx, in, opts...)
}

// 重置密码
func (m *defaultUserService) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.ResetPassword(ctx, in, opts...)
}

// 删除用户
func (m *defaultUserService) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}
