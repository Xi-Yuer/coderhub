package handler

import (
	"net/http"

	"coderhub/api/UserFollow/internal/logic"
	"coderhub/api/UserFollow/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取互相关注列表
func GetMutualFollowsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetMutualFollowsLogic(r.Context(), svcCtx)
		resp, err := l.GetMutualFollows()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
