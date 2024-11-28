// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package handler

import (
	"net/http"

	"coderhub/api/Image/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 健康检查
				Method:  http.MethodGet,
				Path:    "/health",
				Handler: HealthHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/image"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 删除图片
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: DeleteHandler(serverCtx),
			},
			{
				// 获取图片信息
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: GetHandler(serverCtx),
			},
			{
				// 获取用户图片列表
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: ListByUserHandler(serverCtx),
			},
			{
				// 上传图片
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: UploadHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/image"),
	)
}
