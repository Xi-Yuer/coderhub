package handler

import (
	"net/http"

	"coderhub/api/TechSphere/Comment/internal/logic"
	"coderhub/api/TechSphere/Comment/internal/svc"
	"coderhub/api/TechSphere/Comment/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 更新评论点赞数
func UpdateCommentLikeCountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateCommentLikeCountReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateCommentLikeCountLogic(r.Context(), svcCtx)
		resp, err := l.UpdateCommentLikeCount(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
