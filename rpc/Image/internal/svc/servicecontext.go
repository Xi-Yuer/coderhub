package svc

import (
	"coderhub/repository"
	"coderhub/rpc/Image/internal/config"
	"coderhub/shared/storage"
	"fmt"
)

type ServiceContext struct {
	Config          config.Config
	ImageRepository repository.ImageRepository
	Minio           *storage.Minio
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioClient := storage.NewMinio(
		c.Minio.Endpoint,
		c.Minio.AccessKey,
		c.Minio.SecretKey,
		c.Minio.Bucket,
		c.Minio.Region,
		c.Minio.UseSSL,
	)

	err := minioClient.Connect()
	if err != nil {
		panic(fmt.Sprintf("Minio 连接失败: %v", err))
	}

	return &ServiceContext{
		Config:          c,
		ImageRepository: repository.NewImageRepository(storage.NewGorm()),
		Minio:           minioClient,
	}
}
