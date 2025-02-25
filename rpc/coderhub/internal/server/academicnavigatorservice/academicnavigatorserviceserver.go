// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: coderhub.proto

package server

import (
	"context"

	"coderhub/rpc/coderhub/coderhub"
	"coderhub/rpc/coderhub/internal/logic/academicnavigatorservice"
	"coderhub/rpc/coderhub/internal/svc"
)

type AcademicNavigatorServiceServer struct {
	svcCtx *svc.ServiceContext
	coderhub.UnimplementedAcademicNavigatorServiceServer
}

func NewAcademicNavigatorServiceServer(svcCtx *svc.ServiceContext) *AcademicNavigatorServiceServer {
	return &AcademicNavigatorServiceServer{
		svcCtx: svcCtx,
	}
}

// 新增学术导航
func (s *AcademicNavigatorServiceServer) AddAcademicNavigator(ctx context.Context, in *coderhub.AddAcademicNavigatorRequest) (*coderhub.Response, error) {
	l := academicnavigatorservicelogic.NewAddAcademicNavigatorLogic(ctx, s.svcCtx)
	return l.AddAcademicNavigator(in)
}

// 获取学术导航
func (s *AcademicNavigatorServiceServer) GetAcademicNavigator(ctx context.Context, in *coderhub.GetAcademicNavigatorRequest) (*coderhub.GetAcademicNavigatorResponse, error) {
	l := academicnavigatorservicelogic.NewGetAcademicNavigatorLogic(ctx, s.svcCtx)
	return l.GetAcademicNavigator(in)
}

// 删除学术导航
func (s *AcademicNavigatorServiceServer) DeleteAcademicNavigator(ctx context.Context, in *coderhub.DeleteAcademicNavigatorRequest) (*coderhub.Response, error) {
	l := academicnavigatorservicelogic.NewDeleteAcademicNavigatorLogic(ctx, s.svcCtx)
	return l.DeleteAcademicNavigator(in)
}

// 点赞学术导航
func (s *AcademicNavigatorServiceServer) LikeAcademicNavigator(ctx context.Context, in *coderhub.LikeAcademicNavigatorRequest) (*coderhub.Response, error) {
	l := academicnavigatorservicelogic.NewLikeAcademicNavigatorLogic(ctx, s.svcCtx)
	return l.LikeAcademicNavigator(in)
}

// 取消点赞学术导航
func (s *AcademicNavigatorServiceServer) CancelLikeAcademicNavigator(ctx context.Context, in *coderhub.CancelLikeAcademicNavigatorRequest) (*coderhub.Response, error) {
	l := academicnavigatorservicelogic.NewCancelLikeAcademicNavigatorLogic(ctx, s.svcCtx)
	return l.CancelLikeAcademicNavigator(in)
}
