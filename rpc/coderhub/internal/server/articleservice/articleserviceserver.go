// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package server

import (
	"context"

	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/logic/articleservice"
	"coderhub/rpc/coderhub/internal/svc"
)

type ArticleServiceServer struct {
	svcCtx *svc.ServiceContext
	coderhub.UnimplementedArticleServiceServer
}

func NewArticleServiceServer(svcCtx *svc.ServiceContext) *ArticleServiceServer {
	return &ArticleServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *ArticleServiceServer) GetArticle(ctx context.Context, in *coderhub.GetArticleRequest) (*coderhub.GetArticleResponse, error) {
	l := articleservicelogic.NewGetArticleLogic(ctx, s.svcCtx)
	return l.GetArticle(in)
}

func (s *ArticleServiceServer) ListRecommendedArticles(ctx context.Context, in *coderhub.ListRecommendedArticlesRequest) (*coderhub.ListRecommendedArticlesResponse, error) {
	l := articleservicelogic.NewListRecommendedArticlesLogic(ctx, s.svcCtx)
	return l.ListRecommendedArticles(in)
}

func (s *ArticleServiceServer) ListArticles(ctx context.Context, in *coderhub.GetArticlesRequest) (*coderhub.GetArticlesResponse, error) {
	l := articleservicelogic.NewListArticlesLogic(ctx, s.svcCtx)
	return l.ListArticles(in)
}

func (s *ArticleServiceServer) CreateArticle(ctx context.Context, in *coderhub.CreateArticleRequest) (*coderhub.CreateArticleResponse, error) {
	l := articleservicelogic.NewCreateArticleLogic(ctx, s.svcCtx)
	return l.CreateArticle(in)
}

func (s *ArticleServiceServer) UpdateArticle(ctx context.Context, in *coderhub.UpdateArticleRequest) (*coderhub.UpdateArticleResponse, error) {
	l := articleservicelogic.NewUpdateArticleLogic(ctx, s.svcCtx)
	return l.UpdateArticle(in)
}

func (s *ArticleServiceServer) UpdateLikeCount(ctx context.Context, in *coderhub.UpdateLikeCountRequest) (*coderhub.UpdateLikeCountResponse, error) {
	l := articleservicelogic.NewUpdateLikeCountLogic(ctx, s.svcCtx)
	return l.UpdateLikeCount(in)
}

func (s *ArticleServiceServer) DeleteArticle(ctx context.Context, in *coderhub.DeleteArticleRequest) (*coderhub.DeleteArticleResponse, error) {
	l := articleservicelogic.NewDeleteArticleLogic(ctx, s.svcCtx)
	return l.DeleteArticle(in)
}
