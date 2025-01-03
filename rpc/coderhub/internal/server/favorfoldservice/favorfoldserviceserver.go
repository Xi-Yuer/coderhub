// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package server

import (
	"context"

	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/logic/favorfoldservice"
	"coderhub/rpc/coderhub/internal/svc"
)

type FavorFoldServiceServer struct {
	svcCtx *svc.ServiceContext
	coderhub.UnimplementedFavorFoldServiceServer
}

func NewFavorFoldServiceServer(svcCtx *svc.ServiceContext) *FavorFoldServiceServer {
	return &FavorFoldServiceServer{
		svcCtx: svcCtx,
	}
}

// 创建
func (s *FavorFoldServiceServer) CreateFavorFold(ctx context.Context, in *coderhub.CreateFavorFoldRequest) (*coderhub.CreateFavorFoldResponse, error) {
	l := favorfoldservicelogic.NewCreateFavorFoldLogic(ctx, s.svcCtx)
	return l.CreateFavorFold(in)
}

// 删除
func (s *FavorFoldServiceServer) DeleteFavorFold(ctx context.Context, in *coderhub.DeleteFavorFoldRequest) (*coderhub.DeleteFavorFoldResponse, error) {
	l := favorfoldservicelogic.NewDeleteFavorFoldLogic(ctx, s.svcCtx)
	return l.DeleteFavorFold(in)
}

// 更新
func (s *FavorFoldServiceServer) UpdateFavorFold(ctx context.Context, in *coderhub.UpdateFavorFoldRequest) (*coderhub.UpdateFavorFoldResponse, error) {
	l := favorfoldservicelogic.NewUpdateFavorFoldLogic(ctx, s.svcCtx)
	return l.UpdateFavorFold(in)
}

// 获取列表
func (s *FavorFoldServiceServer) GetFavorFoldList(ctx context.Context, in *coderhub.GetFavorFoldListRequest) (*coderhub.GetFavorFoldListResponse, error) {
	l := favorfoldservicelogic.NewGetFavorFoldListLogic(ctx, s.svcCtx)
	return l.GetFavorFoldList(in)
}