// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package server

import (
	"context"

	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/logic/favorservice"
	"coderhub/rpc/coderhub/internal/svc"
)

type FavorServiceServer struct {
	svcCtx *svc.ServiceContext
	coderhub.UnimplementedFavorServiceServer
}

func NewFavorServiceServer(svcCtx *svc.ServiceContext) *FavorServiceServer {
	return &FavorServiceServer{
		svcCtx: svcCtx,
	}
}

// 创建
func (s *FavorServiceServer) CreateFavor(ctx context.Context, in *coderhub.CreateFavorRequest) (*coderhub.CreateFavorResponse, error) {
	l := favorservicelogic.NewCreateFavorLogic(ctx, s.svcCtx)
	return l.CreateFavor(in)
}

// 删除
func (s *FavorServiceServer) DeleteFavor(ctx context.Context, in *coderhub.DeleteFavorRequest) (*coderhub.DeleteFavorResponse, error) {
	l := favorservicelogic.NewDeleteFavorLogic(ctx, s.svcCtx)
	return l.DeleteFavor(in)
}

// 获取列表
func (s *FavorServiceServer) GetFavorList(ctx context.Context, in *coderhub.GetFavorListRequest) (*coderhub.GetFavorListResponse, error) {
	l := favorservicelogic.NewGetFavorListLogic(ctx, s.svcCtx)
	return l.GetFavorList(in)
}
