package svc

import (
	repository "coderhub/repository/Image"
	"coderhub/rpc/Image/internal/config"
	"coderhub/shared/Minio"
	"coderhub/shared/SQL"
	"fmt"
)

type ServiceContext struct {
	Config          config.Config
	ImageRepository repository.ImageRepository
	Minio           *Minio.Minio
}

func NewServiceContext(c config.Config) *ServiceContext {
	minioClient := Minio.NewMinio(
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
		ImageRepository: repository.NewImageRepository(SQL.NewGorm()),
		Minio:           minioClient,
	}
}
