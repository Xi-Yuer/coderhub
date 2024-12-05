package handler

import (
	"net/http"

	"coderhub/api/UserFollow/internal/logic"
	"coderhub/api/UserFollow/internal/svc"
	"coderhub/api/UserFollow/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 检查是否关注
func IsUserFollowedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IsUserFollowedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewIsUserFollowedLogic(r.Context(), svcCtx)
		resp, err := l.IsUserFollowed(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
