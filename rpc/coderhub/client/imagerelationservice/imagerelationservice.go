// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package imagerelationservice

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
	CommentImage                       = coderhub.CommentImage
	CommentUserInfo                    = coderhub.CommentUserInfo
	CreateArticleRequest               = coderhub.CreateArticleRequest
	CreateArticleResponse              = coderhub.CreateArticleResponse
	CreateCommentRequest               = coderhub.CreateCommentRequest
	CreateCommentResponse              = coderhub.CreateCommentResponse
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
	DeleteRequest                      = coderhub.DeleteRequest
	DeleteResponse                     = coderhub.DeleteResponse
	DeleteUserFollowReq                = coderhub.DeleteUserFollowReq
	DeleteUserFollowResp               = coderhub.DeleteUserFollowResp
	DeleteUserRequest                  = coderhub.DeleteUserRequest
	DeleteUserResponse                 = coderhub.DeleteUserResponse
	EntityInfo                         = coderhub.EntityInfo
	GenerateTokenRequest               = coderhub.GenerateTokenRequest
	GenerateTokenResponse              = coderhub.GenerateTokenResponse
	GetAcademicNavigatorRequest        = coderhub.GetAcademicNavigatorRequest
	GetAcademicNavigatorResponse       = coderhub.GetAcademicNavigatorResponse
	GetArticleRequest                  = coderhub.GetArticleRequest
	GetArticleResponse                 = coderhub.GetArticleResponse
	GetCommentRepliesRequest           = coderhub.GetCommentRepliesRequest
	GetCommentRepliesResponse          = coderhub.GetCommentRepliesResponse
	GetCommentRequest                  = coderhub.GetCommentRequest
	GetCommentResponse                 = coderhub.GetCommentResponse
	GetCommentsRequest                 = coderhub.GetCommentsRequest
	GetCommentsResponse                = coderhub.GetCommentsResponse
	GetEntitiesByImageRequest          = coderhub.GetEntitiesByImageRequest
	GetEntitiesByImageResponse         = coderhub.GetEntitiesByImageResponse
	GetImagesByEntityRequest           = coderhub.GetImagesByEntityRequest
	GetImagesByEntityResponse          = coderhub.GetImagesByEntityResponse
	GetMutualFollowsReq                = coderhub.GetMutualFollowsReq
	GetMutualFollowsResp               = coderhub.GetMutualFollowsResp
	GetRequest                         = coderhub.GetRequest
	GetUserFansReq                     = coderhub.GetUserFansReq
	GetUserFansResp                    = coderhub.GetUserFansResp
	GetUserFollowsReq                  = coderhub.GetUserFollowsReq
	GetUserFollowsResp                 = coderhub.GetUserFollowsResp
	GetUserInfoByUsernameRequest       = coderhub.GetUserInfoByUsernameRequest
	GetUserInfoRequest                 = coderhub.GetUserInfoRequest
	GetUserInfoResponse                = coderhub.GetUserInfoResponse
	Image                              = coderhub.Image
	ImageInfo                          = coderhub.ImageInfo
	ImageRelation                      = coderhub.ImageRelation
	IsUserFollowedReq                  = coderhub.IsUserFollowedReq
	IsUserFollowedResp                 = coderhub.IsUserFollowedResp
	LikeAcademicNavigatorRequest       = coderhub.LikeAcademicNavigatorRequest
	ListByUserRequest                  = coderhub.ListByUserRequest
	ListByUserResponse                 = coderhub.ListByUserResponse
	ResetPasswordByLinkRequest         = coderhub.ResetPasswordByLinkRequest
	ResetPasswordByLinkResponse        = coderhub.ResetPasswordByLinkResponse
	ResetPasswordRequest               = coderhub.ResetPasswordRequest
	ResetPasswordResponse              = coderhub.ResetPasswordResponse
	Response                           = coderhub.Response
	UpdateArticleRequest               = coderhub.UpdateArticleRequest
	UpdateArticleResponse              = coderhub.UpdateArticleResponse
	UpdateCommentLikeCountRequest      = coderhub.UpdateCommentLikeCountRequest
	UpdateCommentLikeCountResponse     = coderhub.UpdateCommentLikeCountResponse
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

	ImageRelationService interface {
		// 创建图片关系
		CreateRelation(ctx context.Context, in *CreateRelationRequest, opts ...grpc.CallOption) (*CreateRelationResponse, error)
		// 批量创建图片关系
		BatchCreateRelation(ctx context.Context, in *BatchCreateRelationRequest, opts ...grpc.CallOption) (*BatchCreateRelationResponse, error)
		// 批量删除图片关系
		BatchDeleteRelation(ctx context.Context, in *BatchDeleteRelationRequest, opts ...grpc.CallOption) (*BatchDeleteRelationResponse, error)
		// 批量获取图片关联，根据实体ID列表、实体类型列表获取
		BatchGetImagesByEntity(ctx context.Context, in *BatchGetImagesByEntityRequest, opts ...grpc.CallOption) (*BatchGetImagesByEntityResponse, error)
		// 根据实体ID、实体类型删除图片关系
		DeleteByEntityID(ctx context.Context, in *DeleteByEntityIDRequest, opts ...grpc.CallOption) (*DeleteByEntityIDResponse, error)
		// 获取实体关联的图片列表
		GetImagesByEntity(ctx context.Context, in *GetImagesByEntityRequest, opts ...grpc.CallOption) (*GetImagesByEntityResponse, error)
		// 获取图片关联的实体列表
		GetEntitiesByImage(ctx context.Context, in *GetEntitiesByImageRequest, opts ...grpc.CallOption) (*GetEntitiesByImageResponse, error)
	}

	defaultImageRelationService struct {
		cli zrpc.Client
	}
)

func NewImageRelationService(cli zrpc.Client) ImageRelationService {
	return &defaultImageRelationService{
		cli: cli,
	}
}

// 创建图片关系
func (m *defaultImageRelationService) CreateRelation(ctx context.Context, in *CreateRelationRequest, opts ...grpc.CallOption) (*CreateRelationResponse, error) {
	client := coderhub.NewImageRelationServiceClient(m.cli.Conn())
	return client.CreateRelation(ctx, in, opts...)
}

// 批量创建图片关系
func (m *defaultImageRelationService) BatchCreateRelation(ctx context.Context, in *BatchCreateRelationRequest, opts ...grpc.CallOption) (*BatchCreateRelationResponse, error) {
	client := coderhub.NewImageRelationServiceClient(m.cli.Conn())
	return client.BatchCreateRelation(ctx, in, opts...)
}

// 批量删除图片关系
func (m *defaultImageRelationService) BatchDeleteRelation(ctx context.Context, in *BatchDeleteRelationRequest, opts ...grpc.CallOption) (*BatchDeleteRelationResponse, error) {
	client := coderhub.NewImageRelationServiceClient(m.cli.Conn())
	return client.BatchDeleteRelation(ctx, in, opts...)
}

// 批量获取图片关联，根据实体ID列表、实体类型列表获取
func (m *defaultImageRelationService) BatchGetImagesByEntity(ctx context.Context, in *BatchGetImagesByEntityRequest, opts ...grpc.CallOption) (*BatchGetImagesByEntityResponse, error) {
	client := coderhub.NewImageRelationServiceClient(m.cli.Conn())
	return client.BatchGetImagesByEntity(ctx, in, opts...)
}

// 根据实体ID、实体类型删除图片关系
func (m *defaultImageRelationService) DeleteByEntityID(ctx context.Context, in *DeleteByEntityIDRequest, opts ...grpc.CallOption) (*DeleteByEntityIDResponse, error) {
	client := coderhub.NewImageRelationServiceClient(m.cli.Conn())
	return client.DeleteByEntityID(ctx, in, opts...)
}

// 获取实体关联的图片列表
func (m *defaultImageRelationService) GetImagesByEntity(ctx context.Context, in *GetImagesByEntityRequest, opts ...grpc.CallOption) (*GetImagesByEntityResponse, error) {
	client := coderhub.NewImageRelationServiceClient(m.cli.Conn())
	return client.GetImagesByEntity(ctx, in, opts...)
}

// 获取图片关联的实体列表
func (m *defaultImageRelationService) GetEntitiesByImage(ctx context.Context, in *GetEntitiesByImageRequest, opts ...grpc.CallOption) (*GetEntitiesByImageResponse, error) {
	client := coderhub.NewImageRelationServiceClient(m.cli.Conn())
	return client.GetEntitiesByImage(ctx, in, opts...)
}
