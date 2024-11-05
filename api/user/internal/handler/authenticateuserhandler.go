package handler

import (
	"net/http"

	"coderhub/api/user/internal/logic"
	"coderhub/api/user/internal/svc"
	"coderhub/api/user/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AuthenticateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthenticateUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAuthenticateUserLogic(r.Context(), svcCtx)
		resp, err := l.AuthenticateUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
