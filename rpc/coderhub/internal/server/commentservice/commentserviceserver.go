// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package server

import (
	"context"

	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/logic/commentservice"
	"coderhub/rpc/coderhub/internal/svc"
)

type CommentServiceServer struct {
	svcCtx *svc.ServiceContext
	coderhub.UnimplementedCommentServiceServer
}

func NewCommentServiceServer(svcCtx *svc.ServiceContext) *CommentServiceServer {
	return &CommentServiceServer{
		svcCtx: svcCtx,
	}
}

// 创建评论
func (s *CommentServiceServer) CreateComment(ctx context.Context, in *coderhub.CreateCommentRequest) (*coderhub.CreateCommentResponse, error) {
	l := commentservicelogic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

// 获取评论列表
func (s *CommentServiceServer) GetComments(ctx context.Context, in *coderhub.GetCommentsRequest) (*coderhub.GetCommentsResponse, error) {
	l := commentservicelogic.NewGetCommentsLogic(ctx, s.svcCtx)
	return l.GetComments(in)
}

// 获取某条评论的子评论列表
func (s *CommentServiceServer) GetCommentReplies(ctx context.Context, in *coderhub.GetCommentRepliesRequest) (*coderhub.GetCommentRepliesResponse, error) {
	l := commentservicelogic.NewGetCommentRepliesLogic(ctx, s.svcCtx)
	return l.GetCommentReplies(in)
}

// 更新评论点赞数
func (s *CommentServiceServer) UpdateCommentLikeCount(ctx context.Context, in *coderhub.UpdateCommentLikeCountRequest) (*coderhub.UpdateCommentLikeCountResponse, error) {
	l := commentservicelogic.NewUpdateCommentLikeCountLogic(ctx, s.svcCtx)
	return l.UpdateCommentLikeCount(in)
}

// 获取单个评论详情
func (s *CommentServiceServer) GetComment(ctx context.Context, in *coderhub.GetCommentRequest) (*coderhub.GetCommentResponse, error) {
	l := commentservicelogic.NewGetCommentLogic(ctx, s.svcCtx)
	return l.GetComment(in)
}

// 删除评论
func (s *CommentServiceServer) DeleteComment(ctx context.Context, in *coderhub.DeleteCommentRequest) (*coderhub.DeleteCommentResponse, error) {
	l := commentservicelogic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}
