package config

import (
	"time"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ImageRelationService zrpc.RpcClientConf
	ImageService         zrpc.RpcClientConf
	UserService          zrpc.RpcClientConf
	RabbitMQ             struct {
		Host     string
		Port         string
		Username     string
		Password     string
		MaxRetries   int
		RetryDelay   time.Duration
		ExchangeName string
		QueueName    string
	}
}
