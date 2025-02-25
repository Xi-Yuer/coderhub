// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package favorfoldservice

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
	CreateEmotionRequest               = coderhub.CreateEmotionRequest
	CreateEmotionResponse              = coderhub.CreateEmotionResponse
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
	DeleteEmotionRequest               = coderhub.DeleteEmotionRequest
	DeleteEmotionResponse              = coderhub.DeleteEmotionResponse
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
	Emotion                            = coderhub.Emotion
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
	GetEmotionListRequest              = coderhub.GetEmotionListRequest
	GetEmotionListResponse             = coderhub.GetEmotionListResponse
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

	FavorFoldService interface {
		// 创建
		CreateFavorFold(ctx context.Context, in *CreateFavorFoldRequest, opts ...grpc.CallOption) (*CreateFavorFoldResponse, error)
		// 删除
		DeleteFavorFold(ctx context.Context, in *DeleteFavorFoldRequest, opts ...grpc.CallOption) (*DeleteFavorFoldResponse, error)
		// 更新
		UpdateFavorFold(ctx context.Context, in *UpdateFavorFoldRequest, opts ...grpc.CallOption) (*UpdateFavorFoldResponse, error)
		// 获取列表
		GetFavorFoldList(ctx context.Context, in *GetFavorFoldListRequest, opts ...grpc.CallOption) (*GetFavorFoldListResponse, error)
	}

	defaultFavorFoldService struct {
		cli zrpc.Client
	}
)

func NewFavorFoldService(cli zrpc.Client) FavorFoldService {
	return &defaultFavorFoldService{
		cli: cli,
	}
}

// 创建
func (m *defaultFavorFoldService) CreateFavorFold(ctx context.Context, in *CreateFavorFoldRequest, opts ...grpc.CallOption) (*CreateFavorFoldResponse, error) {
	client := coderhub.NewFavorFoldServiceClient(m.cli.Conn())
	return client.CreateFavorFold(ctx, in, opts...)
}

// 删除
func (m *defaultFavorFoldService) DeleteFavorFold(ctx context.Context, in *DeleteFavorFoldRequest, opts ...grpc.CallOption) (*DeleteFavorFoldResponse, error) {
	client := coderhub.NewFavorFoldServiceClient(m.cli.Conn())
	return client.DeleteFavorFold(ctx, in, opts...)
}

// 更新
func (m *defaultFavorFoldService) UpdateFavorFold(ctx context.Context, in *UpdateFavorFoldRequest, opts ...grpc.CallOption) (*UpdateFavorFoldResponse, error) {
	client := coderhub.NewFavorFoldServiceClient(m.cli.Conn())
	return client.UpdateFavorFold(ctx, in, opts...)
}

// 获取列表
func (m *defaultFavorFoldService) GetFavorFoldList(ctx context.Context, in *GetFavorFoldListRequest, opts ...grpc.CallOption) (*GetFavorFoldListResponse, error) {
	client := coderhub.NewFavorFoldServiceClient(m.cli.Conn())
	return client.GetFavorFoldList(ctx, in, opts...)
}
