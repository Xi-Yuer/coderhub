package MetaData

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func SetUserMetaData(ctx context.Context) context.Context {
	value := ctx.Value("userId")
	md := metadata.New(
		map[string]string{
			"userId": fmt.Sprintf("%v", value), // 将 json.Number 转换为 string
		})
	return metadata.NewOutgoingContext(ctx, md)
}

func GetUserMetaData(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata not found")
	}
	value := md.Get("userId")
	if len(value) == 0 {
		return "", fmt.Errorf("key not found")
	}
	return value[0], nil
}
