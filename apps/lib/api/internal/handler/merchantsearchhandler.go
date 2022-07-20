package handler

import (
	"net/http"

	"goecom/apps/lib/api/internal/logic"
	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func merchantsearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantSearchReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewMerchantsearchLogic(r.Context(), svcCtx)
		resp, err := l.Merchantsearch(&req)
		result.HttpResult(r, w, resp, err)
	}
}
