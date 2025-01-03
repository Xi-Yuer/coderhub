// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package userfollowservice

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

	UserFollowService interface {
		// 创建用户关注关系
		CreateUserFollow(ctx context.Context, in *CreateUserFollowReq, opts ...grpc.CallOption) (*CreateUserFollowResp, error)
		// 删除用户关注关系
		DeleteUserFollow(ctx context.Context, in *DeleteUserFollowReq, opts ...grpc.CallOption) (*DeleteUserFollowResp, error)
		// 获取用户关注列表
		GetUserFollows(ctx context.Context, in *GetUserFollowsReq, opts ...grpc.CallOption) (*GetUserFollowsResp, error)
		// 获取用户粉丝列表
		GetUserFans(ctx context.Context, in *GetUserFansReq, opts ...grpc.CallOption) (*GetUserFansResp, error)
		// 检查是否关注
		IsUserFollowed(ctx context.Context, in *IsUserFollowedReq, opts ...grpc.CallOption) (*IsUserFollowedResp, error)
		// 获取互相关注列表
		GetMutualFollows(ctx context.Context, in *GetMutualFollowsReq, opts ...grpc.CallOption) (*GetMutualFollowsResp, error)
	}

	defaultUserFollowService struct {
		cli zrpc.Client
	}
)

func NewUserFollowService(cli zrpc.Client) UserFollowService {
	return &defaultUserFollowService{
		cli: cli,
	}
}

// 创建用户关注关系
func (m *defaultUserFollowService) CreateUserFollow(ctx context.Context, in *CreateUserFollowReq, opts ...grpc.CallOption) (*CreateUserFollowResp, error) {
	client := coderhub.NewUserFollowServiceClient(m.cli.Conn())
	return client.CreateUserFollow(ctx, in, opts...)
}

// 删除用户关注关系
func (m *defaultUserFollowService) DeleteUserFollow(ctx context.Context, in *DeleteUserFollowReq, opts ...grpc.CallOption) (*DeleteUserFollowResp, error) {
	client := coderhub.NewUserFollowServiceClient(m.cli.Conn())
	return client.DeleteUserFollow(ctx, in, opts...)
}

// 获取用户关注列表
func (m *defaultUserFollowService) GetUserFollows(ctx context.Context, in *GetUserFollowsReq, opts ...grpc.CallOption) (*GetUserFollowsResp, error) {
	client := coderhub.NewUserFollowServiceClient(m.cli.Conn())
	return client.GetUserFollows(ctx, in, opts...)
}

// 获取用户粉丝列表
func (m *defaultUserFollowService) GetUserFans(ctx context.Context, in *GetUserFansReq, opts ...grpc.CallOption) (*GetUserFansResp, error) {
	client := coderhub.NewUserFollowServiceClient(m.cli.Conn())
	return client.GetUserFans(ctx, in, opts...)
}

// 检查是否关注
func (m *defaultUserFollowService) IsUserFollowed(ctx context.Context, in *IsUserFollowedReq, opts ...grpc.CallOption) (*IsUserFollowedResp, error) {
	client := coderhub.NewUserFollowServiceClient(m.cli.Conn())
	return client.IsUserFollowed(ctx, in, opts...)
}

// 获取互相关注列表
func (m *defaultUserFollowService) GetMutualFollows(ctx context.Context, in *GetMutualFollowsReq, opts ...grpc.CallOption) (*GetMutualFollowsResp, error) {
	client := coderhub.NewUserFollowServiceClient(m.cli.Conn())
	return client.GetMutualFollows(ctx, in, opts...)
}
