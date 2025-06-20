package handler

import (
	"net/http"

	"github.com/tiandidoushidali/dex-dodo/backend/app/api/internal/logic"
	"github.com/tiandidoushidali/dex-dodo/backend/app/api/internal/svc"
	"github.com/tiandidoushidali/dex-dodo/backend/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
