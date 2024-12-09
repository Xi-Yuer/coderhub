package handler

import (
	"net/http"

	"coderhub/api/User/internal/logic"
	"coderhub/api/User/internal/svc"
	"coderhub/api/User/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 发送重置密码链接
func SendResetPasswordLinkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendResetPasswordLinkRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendResetPasswordLinkLogic(r.Context(), svcCtx)
		resp, err := l.SendResetPasswordLink(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
