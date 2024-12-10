package handler

import (
	"net/http"

	"coderhub/api/TechSphere/AcademicNavigator/internal/logic"
	"coderhub/api/TechSphere/AcademicNavigator/internal/svc"
	"coderhub/api/TechSphere/AcademicNavigator/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新学术导航点赞数
func UpdateAcademicNavigatorLikeCountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAcademicNavigatorLikeCountReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateAcademicNavigatorLikeCountLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAcademicNavigatorLikeCount(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
