package handler

import (
	"net/http"

	"coderhub/api/User/internal/logic"
	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadUserAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadUserAvatarRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUploadUserAvatarLogic(r.Context(), svcCtx)
		resp, err := l.UploadUserAvatar(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
