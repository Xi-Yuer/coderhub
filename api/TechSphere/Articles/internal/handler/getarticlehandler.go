package handler

import (
	"net/http"

	"coderhub/api/TechSphere/Articles/internal/logic"
	"coderhub/api/TechSphere/Articles/internal/svc"
	"coderhub/api/TechSphere/Articles/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetArticleLogic(r.Context(), svcCtx)
		resp, err := l.GetArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
