package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goecom/apps/lib/api/internal/logic"
	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
)

func storeaddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StoreAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStoreaddLogic(r.Context(), svcCtx)
		resp, err := l.Storeadd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
