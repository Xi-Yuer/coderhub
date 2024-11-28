package handler

import (
	"net/http"

	"coderhub/api/Image/internal/logic"
	"coderhub/api/Image/internal/svc"
	"coderhub/api/Image/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 上传图片
func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
