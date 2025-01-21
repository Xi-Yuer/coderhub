package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserService            zrpc.RpcClientConf
	ImageAuthService       zrpc.RpcClientConf
	ArticlesService        zrpc.RpcClientConf
	AcademicService        zrpc.RpcClientConf
	UserFollowService      zrpc.RpcClientConf
	ImagesService          zrpc.RpcClientConf
	CommentService         zrpc.RpcClientConf
	QuestionBankService    zrpc.RpcClientConf
	FavoriteService        zrpc.RpcClientConf
	FavoriteContentService zrpc.RpcClientConf
	EmotionService         zrpc.RpcClientConf
}
