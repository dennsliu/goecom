package handler

import (
	"net/http"

	"goecom/apps/lib/api/internal/logic"
	"goecom/apps/lib/api/internal/svc"
	"goecom/apps/lib/api/internal/types"
	"goecom/pkg/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func merchantuseraddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MerchantUserAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMerchantuseraddLogic(r.Context(), svcCtx)
		resp, err := l.Merchantuseradd(&req)
		result.HttpResult(r, w, resp, err)
	}
}
