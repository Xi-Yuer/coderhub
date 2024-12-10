package handler

import (
	"net/http"

	"coderhub/api/TechSphere/AcademicNavigator/internal/logic"
	"coderhub/api/TechSphere/AcademicNavigator/internal/svc"
	"coderhub/api/TechSphere/AcademicNavigator/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 点赞学术导航
func PostAcademicNavigatorLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostAcademicNavigatorLikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPostAcademicNavigatorLikeLogic(r.Context(), svcCtx)
		resp, err := l.PostAcademicNavigatorLike(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
