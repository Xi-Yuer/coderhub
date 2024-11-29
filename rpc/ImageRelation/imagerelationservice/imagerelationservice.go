// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: imageRelation.proto

package imagerelationservice

import (
	"context"

	"coderhub/rpc/ImageRelation/imageRelation"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchCreateRelationRequest  = imageRelation.BatchCreateRelationRequest
	BatchCreateRelationResponse = imageRelation.BatchCreateRelationResponse
	BatchDeleteRelationRequest  = imageRelation.BatchDeleteRelationRequest
	BatchDeleteRelationResponse = imageRelation.BatchDeleteRelationResponse
	CreateRelationRequest       = imageRelation.CreateRelationRequest
	CreateRelationResponse      = imageRelation.CreateRelationResponse
	DeleteByEntityIDRequest     = imageRelation.DeleteByEntityIDRequest
	DeleteByEntityIDResponse    = imageRelation.DeleteByEntityIDResponse
	EntityInfo                  = imageRelation.EntityInfo
	GetEntitiesByImageRequest   = imageRelation.GetEntitiesByImageRequest
	GetEntitiesByImageResponse  = imageRelation.GetEntitiesByImageResponse
	GetImagesByEntityRequest    = imageRelation.GetImagesByEntityRequest
	GetImagesByEntityResponse   = imageRelation.GetImagesByEntityResponse
	ImageInfo                   = imageRelation.ImageInfo
	ImageRelation               = imageRelation.ImageRelation

	ImageRelationService interface {
		// 创建图片关系
		CreateRelation(ctx context.Context, in *CreateRelationRequest, opts ...grpc.CallOption) (*CreateRelationResponse, error)
		// 批量创建图片关系
		BatchCreateRelation(ctx context.Context, in *BatchCreateRelationRequest, opts ...grpc.CallOption) (*BatchCreateRelationResponse, error)
		// 批量删除图片关系
		BatchDeleteRelation(ctx context.Context, in *BatchDeleteRelationRequest, opts ...grpc.CallOption) (*BatchDeleteRelationResponse, error)
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
	client := imageRelation.NewImageRelationServiceClient(m.cli.Conn())
	return client.CreateRelation(ctx, in, opts...)
}

// 批量创建图片关系
func (m *defaultImageRelationService) BatchCreateRelation(ctx context.Context, in *BatchCreateRelationRequest, opts ...grpc.CallOption) (*BatchCreateRelationResponse, error) {
	client := imageRelation.NewImageRelationServiceClient(m.cli.Conn())
	return client.BatchCreateRelation(ctx, in, opts...)
}

// 批量删除图片关系
func (m *defaultImageRelationService) BatchDeleteRelation(ctx context.Context, in *BatchDeleteRelationRequest, opts ...grpc.CallOption) (*BatchDeleteRelationResponse, error) {
	client := imageRelation.NewImageRelationServiceClient(m.cli.Conn())
	return client.BatchDeleteRelation(ctx, in, opts...)
}

// 根据实体ID、实体类型删除图片关系
func (m *defaultImageRelationService) DeleteByEntityID(ctx context.Context, in *DeleteByEntityIDRequest, opts ...grpc.CallOption) (*DeleteByEntityIDResponse, error) {
	client := imageRelation.NewImageRelationServiceClient(m.cli.Conn())
	return client.DeleteByEntityID(ctx, in, opts...)
}

// 获取实体关联的图片列表
func (m *defaultImageRelationService) GetImagesByEntity(ctx context.Context, in *GetImagesByEntityRequest, opts ...grpc.CallOption) (*GetImagesByEntityResponse, error) {
	client := imageRelation.NewImageRelationServiceClient(m.cli.Conn())
	return client.GetImagesByEntity(ctx, in, opts...)
}

// 获取图片关联的实体列表
func (m *defaultImageRelationService) GetEntitiesByImage(ctx context.Context, in *GetEntitiesByImageRequest, opts ...grpc.CallOption) (*GetEntitiesByImageResponse, error) {
	client := imageRelation.NewImageRelationServiceClient(m.cli.Conn())
	return client.GetEntitiesByImage(ctx, in, opts...)
}
