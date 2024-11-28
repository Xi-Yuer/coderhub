package handler

import (
	"net/http"

	"coderhub/api/Image/internal/logic"
	"coderhub/api/Image/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 健康检查
func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHealthLogic(r.Context(), svcCtx)
		resp, err := l.Health()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
