package handler

import (
	"net/http"

	"coderhub/api/User/internal/logic"
	"coderhub/api/User/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadUserAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadUserAvatarLogic(r.Context(), r, svcCtx)
		resp, err := l.UploadUserAvatar()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
