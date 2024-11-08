package metaData

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func SetMetaData(ctx context.Context, key string, value interface{}) context.Context {
	md := metadata.New(
		map[string]string{
			key: fmt.Sprintf("%v", value), // 将 json.Number 转换为 string
		})
	return metadata.NewOutgoingContext(ctx, md)
}

func GetMetaData(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata not found")
	}
	value := md.Get(key)
	if len(value) == 0 {
		return "", fmt.Errorf("key not found")
	}
	return value[0], nil
}
