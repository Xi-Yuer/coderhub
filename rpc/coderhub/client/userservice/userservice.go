// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package userservice

import (
	"context"

	"coderhub/rpc/coderhub/coderhub"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AcademicNavigator                  = coderhub.AcademicNavigator
	AddAcademicNavigatorRequest        = coderhub.AddAcademicNavigatorRequest
	Article                            = coderhub.Article
	AuthorizeRequest                   = coderhub.AuthorizeRequest
	AuthorizeResponse                  = coderhub.AuthorizeResponse
	BatchCreateRelationRequest         = coderhub.BatchCreateRelationRequest
	BatchCreateRelationResponse        = coderhub.BatchCreateRelationResponse
	BatchDeleteRelationRequest         = coderhub.BatchDeleteRelationRequest
	BatchDeleteRelationResponse        = coderhub.BatchDeleteRelationResponse
	BatchGetImagesByEntityRequest      = coderhub.BatchGetImagesByEntityRequest
	BatchGetImagesByEntityResponse     = coderhub.BatchGetImagesByEntityResponse
	BatchGetRequest                    = coderhub.BatchGetRequest
	BatchGetResponse                   = coderhub.BatchGetResponse
	BatchGetUserByIDRequest            = coderhub.BatchGetUserByIDRequest
	BatchGetUserByIDResponse           = coderhub.BatchGetUserByIDResponse
	CancelLikeAcademicNavigatorRequest = coderhub.CancelLikeAcademicNavigatorRequest
	ChangePasswordRequest              = coderhub.ChangePasswordRequest
	ChangePasswordResponse             = coderhub.ChangePasswordResponse
	CheckUserExistsRequest             = coderhub.CheckUserExistsRequest
	CheckUserExistsResponse            = coderhub.CheckUserExistsResponse
	Comment                            = coderhub.Comment
	CreateArticleRequest               = coderhub.CreateArticleRequest
	CreateArticleResponse              = coderhub.CreateArticleResponse
	CreateCommentRequest               = coderhub.CreateCommentRequest
	CreateCommentResponse              = coderhub.CreateCommentResponse
	CreateFavorFoldRequest             = coderhub.CreateFavorFoldRequest
	CreateFavorFoldResponse            = coderhub.CreateFavorFoldResponse
	CreateFavorRequest                 = coderhub.CreateFavorRequest
	CreateFavorResponse                = coderhub.CreateFavorResponse
	CreateQuestionBankRequest          = coderhub.CreateQuestionBankRequest
	CreateQuestionBankResponse         = coderhub.CreateQuestionBankResponse
	CreateQuestionRequest              = coderhub.CreateQuestionRequest
	CreateQuestionResponse             = coderhub.CreateQuestionResponse
	CreateRelationRequest              = coderhub.CreateRelationRequest
	CreateRelationResponse             = coderhub.CreateRelationResponse
	CreateUserFollowReq                = coderhub.CreateUserFollowReq
	CreateUserFollowResp               = coderhub.CreateUserFollowResp
	CreateUserRequest                  = coderhub.CreateUserRequest
	CreateUserResponse                 = coderhub.CreateUserResponse
	DeleteAcademicNavigatorRequest     = coderhub.DeleteAcademicNavigatorRequest
	DeleteArticleRequest               = coderhub.DeleteArticleRequest
	DeleteArticleResponse              = coderhub.DeleteArticleResponse
	DeleteByEntityIDRequest            = coderhub.DeleteByEntityIDRequest
	DeleteByEntityIDResponse           = coderhub.DeleteByEntityIDResponse
	DeleteCommentRequest               = coderhub.DeleteCommentRequest
	DeleteCommentResponse              = coderhub.DeleteCommentResponse
	DeleteFavorFoldRequest             = coderhub.DeleteFavorFoldRequest
	DeleteFavorFoldResponse            = coderhub.DeleteFavorFoldResponse
	DeleteFavorRequest                 = coderhub.DeleteFavorRequest
	DeleteFavorResponse                = coderhub.DeleteFavorResponse
	DeleteQuestionBankRequest          = coderhub.DeleteQuestionBankRequest
	DeleteQuestionBankResponse         = coderhub.DeleteQuestionBankResponse
	DeleteQuestionRequest              = coderhub.DeleteQuestionRequest
	DeleteQuestionResponse             = coderhub.DeleteQuestionResponse
	DeleteRequest                      = coderhub.DeleteRequest
	DeleteResponse                     = coderhub.DeleteResponse
	DeleteUserFollowReq                = coderhub.DeleteUserFollowReq
	DeleteUserFollowResp               = coderhub.DeleteUserFollowResp
	DeleteUserRequest                  = coderhub.DeleteUserRequest
	DeleteUserResponse                 = coderhub.DeleteUserResponse
	EntityInfo                         = coderhub.EntityInfo
	Favor                              = coderhub.Favor
	FavorFold                          = coderhub.FavorFold
	FavorPreview                       = coderhub.FavorPreview
	GenerateTokenRequest               = coderhub.GenerateTokenRequest
	GenerateTokenResponse              = coderhub.GenerateTokenResponse
	GetAcademicNavigatorRequest        = coderhub.GetAcademicNavigatorRequest
	GetAcademicNavigatorResponse       = coderhub.GetAcademicNavigatorResponse
	GetArticleRequest                  = coderhub.GetArticleRequest
	GetArticleResponse                 = coderhub.GetArticleResponse
	GetArticlesRequest                 = coderhub.GetArticlesRequest
	GetArticlesResponse                = coderhub.GetArticlesResponse
	GetCommentRepliesRequest           = coderhub.GetCommentRepliesRequest
	GetCommentRepliesResponse          = coderhub.GetCommentRepliesResponse
	GetCommentRequest                  = coderhub.GetCommentRequest
	GetCommentResponse                 = coderhub.GetCommentResponse
	GetCommentsRequest                 = coderhub.GetCommentsRequest
	GetCommentsResponse                = coderhub.GetCommentsResponse
	GetEntitiesByImageRequest          = coderhub.GetEntitiesByImageRequest
	GetEntitiesByImageResponse         = coderhub.GetEntitiesByImageResponse
	GetFavorFoldListRequest            = coderhub.GetFavorFoldListRequest
	GetFavorFoldListResponse           = coderhub.GetFavorFoldListResponse
	GetFavorListRequest                = coderhub.GetFavorListRequest
	GetFavorListResponse               = coderhub.GetFavorListResponse
	GetImagesByEntityRequest           = coderhub.GetImagesByEntityRequest
	GetImagesByEntityResponse          = coderhub.GetImagesByEntityResponse
	GetMutualFollowsReq                = coderhub.GetMutualFollowsReq
	GetMutualFollowsResp               = coderhub.GetMutualFollowsResp
	GetQuestionBankListRequest         = coderhub.GetQuestionBankListRequest
	GetQuestionBankListResponse        = coderhub.GetQuestionBankListResponse
	GetQuestionRequest                 = coderhub.GetQuestionRequest
	GetQuestionResponse                = coderhub.GetQuestionResponse
	GetQuestionTreeRequest             = coderhub.GetQuestionTreeRequest
	GetQuestionTreeResponse            = coderhub.GetQuestionTreeResponse
	GetRequest                         = coderhub.GetRequest
	GetUserFansReq                     = coderhub.GetUserFansReq
	GetUserFansResp                    = coderhub.GetUserFansResp
	GetUserFollowsReq                  = coderhub.GetUserFollowsReq
	GetUserFollowsResp                 = coderhub.GetUserFollowsResp
	GetUserInfoByUsernameRequest       = coderhub.GetUserInfoByUsernameRequest
	GetUserInfoRequest                 = coderhub.GetUserInfoRequest
	Image                              = coderhub.Image
	ImageInfo                          = coderhub.ImageInfo
	ImageRelation                      = coderhub.ImageRelation
	IsUserFollowedReq                  = coderhub.IsUserFollowedReq
	IsUserFollowedResp                 = coderhub.IsUserFollowedResp
	LikeAcademicNavigatorRequest       = coderhub.LikeAcademicNavigatorRequest
	ListByUserRequest                  = coderhub.ListByUserRequest
	ListByUserResponse                 = coderhub.ListByUserResponse
	ListRecommendedArticlesRequest     = coderhub.ListRecommendedArticlesRequest
	ListRecommendedArticlesResponse    = coderhub.ListRecommendedArticlesResponse
	QuestionBank                       = coderhub.QuestionBank
	QuestionTreeNode                   = coderhub.QuestionTreeNode
	ResetPasswordByLinkRequest         = coderhub.ResetPasswordByLinkRequest
	ResetPasswordByLinkResponse        = coderhub.ResetPasswordByLinkResponse
	ResetPasswordRequest               = coderhub.ResetPasswordRequest
	ResetPasswordResponse              = coderhub.ResetPasswordResponse
	Response                           = coderhub.Response
	UpdateArticleRequest               = coderhub.UpdateArticleRequest
	UpdateArticleResponse              = coderhub.UpdateArticleResponse
	UpdateCommentLikeCountRequest      = coderhub.UpdateCommentLikeCountRequest
	UpdateCommentLikeCountResponse     = coderhub.UpdateCommentLikeCountResponse
	UpdateFavorFoldRequest             = coderhub.UpdateFavorFoldRequest
	UpdateFavorFoldResponse            = coderhub.UpdateFavorFoldResponse
	UpdateLikeCountRequest             = coderhub.UpdateLikeCountRequest
	UpdateLikeCountResponse            = coderhub.UpdateLikeCountResponse
	UpdateUserInfoRequest              = coderhub.UpdateUserInfoRequest
	UpdateUserInfoResponse             = coderhub.UpdateUserInfoResponse
	UploadAvatarRequest                = coderhub.UploadAvatarRequest
	UploadAvatarResponse               = coderhub.UploadAvatarResponse
	UploadImageInfo                    = coderhub.UploadImageInfo
	UploadRequest                      = coderhub.UploadRequest
	UserFollowInfo                     = coderhub.UserFollowInfo
	UserInfo                           = coderhub.UserInfo

	UserService interface {
		// 授权
		Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
		// 检查用户是否存在
		CheckUserExists(ctx context.Context, in *CheckUserExistsRequest, opts ...grpc.CallOption) (*CheckUserExistsResponse, error)
		// 创建用户
		CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
		// 获取用户信息
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error)
		GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*UserInfo, error)
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
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.Authorize(ctx, in, opts...)
}

// 检查用户是否存在
func (m *defaultUserService) CheckUserExists(ctx context.Context, in *CheckUserExistsRequest, opts ...grpc.CallOption) (*CheckUserExistsResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.CheckUserExists(ctx, in, opts...)
}

// 创建用户
func (m *defaultUserService) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

// 获取用户信息
func (m *defaultUserService) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserService) GetUserInfoByUsername(ctx context.Context, in *GetUserInfoByUsernameRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfoByUsername(ctx, in, opts...)
}

// 批量获取用户信息
func (m *defaultUserService) BatchGetUserByID(ctx context.Context, in *BatchGetUserByIDRequest, opts ...grpc.CallOption) (*BatchGetUserByIDResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.BatchGetUserByID(ctx, in, opts...)
}

// 更新用户信息
func (m *defaultUserService) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUserInfo(ctx, in, opts...)
}

// 上传用户头像
func (m *defaultUserService) UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.UploadAvatar(ctx, in, opts...)
}

// 修改密码
func (m *defaultUserService) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.ChangePassword(ctx, in, opts...)
}

// 重置密码, 通过邮箱发送重置密码链接
func (m *defaultUserService) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.ResetPassword(ctx, in, opts...)
}

// 通过链接重置密码
func (m *defaultUserService) ResetPasswordByLink(ctx context.Context, in *ResetPasswordByLinkRequest, opts ...grpc.CallOption) (*ResetPasswordByLinkResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.ResetPasswordByLink(ctx, in, opts...)
}

// 删除用户
func (m *defaultUserService) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	client := coderhub.NewUserServiceClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}
