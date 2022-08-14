package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goecom/apps/lib/api/internal/logic"
	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
)

func storeupdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StoreUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStoreupdateLogic(r.Context(), svcCtx)
		resp, err := l.Storeupdate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
