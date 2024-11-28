// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: image.proto

package imageclient

import (
	"context"

	"coderhub/rpc/Image/image"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DeleteRequest      = image.DeleteRequest
	DeleteResponse     = image.DeleteResponse
	GetRequest         = image.GetRequest
	ImageInfo          = image.ImageInfo
	ListByUserRequest  = image.ListByUserRequest
	ListByUserResponse = image.ListByUserResponse
	UploadRequest      = image.UploadRequest

	Image interface {
		// 上传图片
		Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*ImageInfo, error)
		// 删除图片
		Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
		// 获取图片信息
		Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ImageInfo, error)
		// 获取用户图片列表
		ListByUser(ctx context.Context, in *ListByUserRequest, opts ...grpc.CallOption) (*ListByUserResponse, error)
	}

	defaultImage struct {
		cli zrpc.Client
	}
)

func NewImage(cli zrpc.Client) Image {
	return &defaultImage{
		cli: cli,
	}
}

// 上传图片
func (m *defaultImage) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*ImageInfo, error) {
	client := image.NewImageClient(m.cli.Conn())
	return client.Upload(ctx, in, opts...)
}

// 删除图片
func (m *defaultImage) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	client := image.NewImageClient(m.cli.Conn())
	return client.Delete(ctx, in, opts...)
}

// 获取图片信息
func (m *defaultImage) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*ImageInfo, error) {
	client := image.NewImageClient(m.cli.Conn())
	return client.Get(ctx, in, opts...)
}

// 获取用户图片列表
func (m *defaultImage) ListByUser(ctx context.Context, in *ListByUserRequest, opts ...grpc.CallOption) (*ListByUserResponse, error) {
	client := image.NewImageClient(m.cli.Conn())
	return client.ListByUser(ctx, in, opts...)
}