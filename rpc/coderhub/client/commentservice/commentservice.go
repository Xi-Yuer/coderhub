// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package commentservice

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

	CommentService interface {
		// 创建评论
		CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
		// 获取评论列表
		GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error)
		// 获取某条评论的子评论列表
		GetCommentReplies(ctx context.Context, in *GetCommentRepliesRequest, opts ...grpc.CallOption) (*GetCommentRepliesResponse, error)
		// 更新评论点赞数
		UpdateCommentLikeCount(ctx context.Context, in *UpdateCommentLikeCountRequest, opts ...grpc.CallOption) (*UpdateCommentLikeCountResponse, error)
		// 获取单个评论详情
		GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
		// 删除评论
		DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
	}

	defaultCommentService struct {
		cli zrpc.Client
	}
)

func NewCommentService(cli zrpc.Client) CommentService {
	return &defaultCommentService{
		cli: cli,
	}
}

// 创建评论
func (m *defaultCommentService) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	client := coderhub.NewCommentServiceClient(m.cli.Conn())
	return client.CreateComment(ctx, in, opts...)
}

// 获取评论列表
func (m *defaultCommentService) GetComments(ctx context.Context, in *GetCommentsRequest, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	client := coderhub.NewCommentServiceClient(m.cli.Conn())
	return client.GetComments(ctx, in, opts...)
}

// 获取某条评论的子评论列表
func (m *defaultCommentService) GetCommentReplies(ctx context.Context, in *GetCommentRepliesRequest, opts ...grpc.CallOption) (*GetCommentRepliesResponse, error) {
	client := coderhub.NewCommentServiceClient(m.cli.Conn())
	return client.GetCommentReplies(ctx, in, opts...)
}

// 更新评论点赞数
func (m *defaultCommentService) UpdateCommentLikeCount(ctx context.Context, in *UpdateCommentLikeCountRequest, opts ...grpc.CallOption) (*UpdateCommentLikeCountResponse, error) {
	client := coderhub.NewCommentServiceClient(m.cli.Conn())
	return client.UpdateCommentLikeCount(ctx, in, opts...)
}

// 获取单个评论详情
func (m *defaultCommentService) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	client := coderhub.NewCommentServiceClient(m.cli.Conn())
	return client.GetComment(ctx, in, opts...)
}

// 删除评论
func (m *defaultCommentService) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	client := coderhub.NewCommentServiceClient(m.cli.Conn())
	return client.DeleteComment(ctx, in, opts...)
}
