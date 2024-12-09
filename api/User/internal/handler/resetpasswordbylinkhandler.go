package handler

import (
	"net/http"

	"coderhub/api/User/internal/logic"
	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 通过链接重置密码
func ResetPasswordByLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetPasswordByLinkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewResetPasswordByLinkLogic(r.Context(), svcCtx)
		resp, err := l.ResetPasswordByLink(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
