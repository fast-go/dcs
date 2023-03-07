package handler

import (
	"net/http"

	"dcs/api/frontend/internal/logic/product"
	"dcs/api/frontend/internal/svc"
	"dcs/api/frontend/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductGetDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := product.NewProductLogic(r.Context(), ctx)
		resp, err := l.Product(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}




