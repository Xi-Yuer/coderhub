package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	ImageRelationService zrpc.RpcClientConf
	ImageService         zrpc.RpcClientConf
	UserService          zrpc.RpcClientConf
}
