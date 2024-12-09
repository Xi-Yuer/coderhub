// goctl rpc protoc user.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.3
// source: user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_Authorize_FullMethodName             = "/user.UserService/Authorize"
	UserService_CheckUserExists_FullMethodName       = "/user.UserService/CheckUserExists"
	UserService_CreateUser_FullMethodName            = "/user.UserService/CreateUser"
	UserService_GetUserInfo_FullMethodName           = "/user.UserService/GetUserInfo"
	UserService_GetUserInfoByUsername_FullMethodName = "/user.UserService/GetUserInfoByUsername"
	UserService_BatchGetUserByID_FullMethodName      = "/user.UserService/BatchGetUserByID"
	UserService_UpdateUserInfo_FullMethodName        = "/user.UserService/UpdateUserInfo"
	UserService_UploadAvatar_FullMethodName          = "/user.UserService/UploadAvatar"
	UserService_ChangePassword_FullMethodName        = "/user.UserService/ChangePassword"
	UserService_ResetPassword_FullMethodName         = "/user.UserService/ResetPassword"
	UserService_ResetPasswordByLink_FullMethodName   = "/user.UserService/ResetPasswordByLink"
	UserService_DeleteUser_FullMethodName            = "/user.UserService/DeleteUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 授权
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	// 检查用户是否存在
	CheckUserExists(ctx context.Context, in *CheckUserExistsRequest, opts ...grpc.CallOption) (*CheckUserExistsResponse, error)
	// 创建用户
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	// 获取用户信息
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	// 批量获取用户信息
	BatchGetUserByID(ctx context.Context, in *BatchGetUserByIDRequest, opts ...grpc.CallOption) (*BatchGetUserByIDResponse, error)
	// 更新用户信息
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	// 上传用户头像
	UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error)
	// 修改密码
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	// 重置密码, 通过邮箱发送重置密码链接
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	// 通过链接重置密码
	ResetPasswordByLink(ctx context.Context, in *ResetPasswordByLinkRequest, opts ...grpc.CallOption) (*ResetPasswordByLinkResponse, error)
	// 删除用户
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, UserService_Authorize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckUserExists(ctx context.Context, in *CheckUserExistsRequest, opts ...grpc.CallOption) (*CheckUserExistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckUserExistsResponse)
	err := c.cc.Invoke(ctx, UserService_CheckUserExists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserInfoByUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BatchGetUserByID(ctx context.Context, in *BatchGetUserByIDRequest, opts ...grpc.CallOption) (*BatchGetUserByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetUserByIDResponse)
	err := c.cc.Invoke(ctx, UserService_BatchGetUserByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUserInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UploadAvatarResponse)
	err := c.cc.Invoke(ctx, UserService_UploadAvatar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, UserService_ChangePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, UserService_ResetPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPasswordByLink(ctx context.Context, in *ResetPasswordByLinkRequest, opts ...grpc.CallOption) (*ResetPasswordByLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResetPasswordByLinkResponse)
	err := c.cc.Invoke(ctx, UserService_ResetPasswordByLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	// 授权
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	// 检查用户是否存在
	CheckUserExists(context.Context, *CheckUserExistsRequest) (*CheckUserExistsResponse, error)
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	// 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	GetUserInfoByUsername(context.Context, *GetUserInfoByUsernameRequest) (*GetUserInfoResponse, error)
	// 批量获取用户信息
	BatchGetUserByID(context.Context, *BatchGetUserByIDRequest) (*BatchGetUserByIDResponse, error)
	// 更新用户信息
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	// 上传用户头像
	UploadAvatar(context.Context, *UploadAvatarRequest) (*UploadAvatarResponse, error)
	// 修改密码
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	// 重置密码, 通过邮箱发送重置密码链接
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	// 通过链接重置密码
	ResetPasswordByLink(context.Context, *ResetPasswordByLinkRequest) (*ResetPasswordByLinkResponse, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedUserServiceServer) CheckUserExists(context.Context, *CheckUserExistsRequest) (*CheckUserExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserExists not implemented")
}
func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfoByUsername(context.Context, *GetUserInfoByUsernameRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfoByUsername not implemented")
}
func (UnimplementedUserServiceServer) BatchGetUserByID(context.Context, *BatchGetUserByIDRequest) (*BatchGetUserByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetUserByID not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserServiceServer) UploadAvatar(context.Context, *UploadAvatarRequest) (*UploadAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedUserServiceServer) ResetPasswordByLink(context.Context, *ResetPasswordByLinkRequest) (*ResetPasswordByLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPasswordByLink not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckUserExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckUserExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckUserExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckUserExists(ctx, req.(*CheckUserExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfoByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfoByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfoByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfoByUsername(ctx, req.(*GetUserInfoByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BatchGetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BatchGetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_BatchGetUserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BatchGetUserByID(ctx, req.(*BatchGetUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UploadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UploadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UploadAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UploadAvatar(ctx, req.(*UploadAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ChangePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPasswordByLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordByLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPasswordByLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ResetPasswordByLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPasswordByLink(ctx, req.(*ResetPasswordByLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _UserService_Authorize_Handler,
		},
		{
			MethodName: "CheckUserExists",
			Handler:    _UserService_CheckUserExists_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfoByUsername",
			Handler:    _UserService_GetUserInfoByUsername_Handler,
		},
		{
			MethodName: "BatchGetUserByID",
			Handler:    _UserService_BatchGetUserByID_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "UploadAvatar",
			Handler:    _UserService_UploadAvatar_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "ResetPasswordByLink",
			Handler:    _UserService_ResetPasswordByLink_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
